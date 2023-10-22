package podcast

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/h2non/filetype"
	"github.com/rs/zerolog"
	"github.com/upamune/radicaster/metadata"
)

type Podcast struct {
	Title       string
	Link        string
	Description string
	PublishedAt *time.Time
	ImageURL    string

	Episodes []Episode
}

type Episode struct {
	Title         string
	Description   string
	PublishedAt   *time.Time
	URL           string
	LengthInBytes int64
	ImageURL      string
	PodcastTitle  string
}

type Podcaster struct {
	logger zerolog.Logger

	baseURL   string
	targetDir string

	title       string
	link        string
	description string
	publishedAt *time.Time
	imageURL    string

	mu      *sync.RWMutex
	feedMap map[string]string
}

func NewPodcaster(
	logger zerolog.Logger,
	baseURL string,
	targetDir string,
	title string,
	link string,
	description string,
	publishedAt *time.Time,
	imageURL string,
) *Podcaster {
	p := &Podcaster{
		logger:      logger,
		baseURL:     baseURL,
		targetDir:   targetDir,
		title:       title,
		link:        link,
		description: description,
		publishedAt: publishedAt,
		imageURL:    imageURL,
		mu:          &sync.RWMutex{},
	}
	return p
}

func (p *Podcaster) GetDefaultFeed() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.feedMap[""]
}

func (p *Podcaster) GetFeed(path string) (string, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	feed, ok := p.feedMap[path]
	return feed, ok
}

func sortEpisodesByPublishedAtDesc(episodes []Episode) {
	slices.SortStableFunc(episodes, func(a, b Episode) int {
		if a.PublishedAt.Unix() == b.PublishedAt.Unix() {
			return 0
		}
		// NOTE: 降順にしたいので逆にしている
		if a.PublishedAt.Unix() < b.PublishedAt.Unix() {
			return 1
		}
		return -1
	})
}

func (p *Podcaster) Sync() error {
	p.logger.Info().Msg("Podcaster.Sync started")
	defer func() {
		p.logger.Info().Msg("Podcaster.Sync ended")
	}()

	var (
		allEpisodes         []Episode
		pathGroupedEpisodes = make(map[string][]Episode)
	)
	p.logger.Info().Str("target_dir", p.targetDir).Msg("filepath.Walk is starting")
	if err := filepath.Walk(p.targetDir, func(fpath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		p.logger.Info().Str("path", fpath).Err(err).Msg("found a target file")

		if !p.isAudioFile(fpath) {
			p.logger.Info().
				Str("path", fpath).
				Msg("skip because the file is not audio file")
			return nil
		}

		baseName := filepath.Base(fpath)

		u, err := url.Parse(p.baseURL)
		if err != nil {
			return fmt.Errorf("failed to parse baseURL(%s): %w", p.baseURL, err)
		}
		u.Path = path.Join(u.Path, "static", baseName)

		stat, err := os.Stat(fpath)
		if err != nil {
			return err
		}
		ep := Episode{
			Title:         fpath,
			URL:           u.String(),
			LengthInBytes: stat.Size(),
		}
		if ss := strings.Split(baseName, "_"); len(ss) > 1 {
			ep.Title = ss[0]
			if startedAt, _ := time.Parse("200601021504", strings.TrimSuffix(ss[1], filepath.Ext(ss[1]))); err == nil {
				ep.PublishedAt = &startedAt
			}
		}
		if ep.PublishedAt == nil {
			now := time.Now()
			ep.PublishedAt = &now
		}

		var podcastPath string
		// NOTE: メタデータがあればそれで全て上書きする
		if md, err := metadata.ReadByAudioFilePath(fpath); err == nil {
			ep.Title = md.Title
			ep.Description = md.Description
			ep.PublishedAt = &md.PublishedAt
			ep.ImageURL = md.ImageURL
			ep.PodcastTitle = md.PodcastTitle

			// NOTE: `/ann` のような設定を `ann` と同値にしてあげる
			podcastPath = strings.ToLower(strings.TrimPrefix(md.Path, "/"))
			if md.ZenrokuMode {
				podcastPath = path.Join("zenroku", podcastPath)
			}
		}

		allEpisodes = append(allEpisodes, ep)
		pathGroupedEpisodes[podcastPath] = append(pathGroupedEpisodes[podcastPath], ep)

		return nil
	}); err != nil {
		return err
	}

	feedMap := make(map[string]string)

	encodePodcastToXML := func(podcast *Podcast) (string, error) {
		buf := bytes.NewBuffer(nil)
		p.logger.Info().Msg("encodeXML is starting")
		if err := encodeXML(buf, podcast); err != nil {
			return "", errors.Wrap(err, "failed to encodeXM")
		}
		return buf.String(), nil
	}
	for path, episodes := range pathGroupedEpisodes {
		path, episodes := path, episodes

		if len(episodes) == 0 {
			continue
		}

		sortEpisodesByPublishedAtDesc(episodes)
		latestEpisode := episodes[0]

		p.logger.Debug().
			Str("path", path).
			Int("episodes_count", len(episodes)).
			Str("title", latestEpisode.Title).
			Time("published_at", *latestEpisode.PublishedAt).
			Msg("latestEpisode is found")

		podcastTitle := p.title
		if latestEpisode.PodcastTitle != "" {
			podcastTitle = latestEpisode.PodcastTitle
		}
		podcast := &Podcast{
			Title:       podcastTitle,
			Link:        p.link,
			Description: p.description,
			PublishedAt: p.publishedAt,
			ImageURL:    latestEpisode.ImageURL,
		}

		// NOTE: デフォルトパス(= "")の場合はデフォルト設定にする
		if path == "" {
			podcast = &Podcast{
				Title:       p.title,
				Link:        p.link,
				Description: p.description,
				PublishedAt: p.publishedAt,
				ImageURL:    p.imageURL,
			}
		}

		podcast.Episodes = episodes
		feed, err := encodePodcastToXML(podcast)
		if err != nil {
			p.logger.Err(err).
				Str("path", path).
				Msg("failed to encodeXML")
			return errors.Wrapf(err, "path=%s", path)
		}
		feedMap[path] = feed
	}

	sortEpisodesByPublishedAtDesc(allEpisodes)
	feed, err := encodePodcastToXML(
		&Podcast{
			Title:       fmt.Sprintf("%s(ALL)", p.title),
			Link:        p.link,
			Description: p.description,
			PublishedAt: p.publishedAt,
			ImageURL:    p.imageURL,
			Episodes:    allEpisodes,
		},
	)
	if err != nil {
		return errors.Wrap(err, "all episodes")
	}
	p.logger.Trace().Str("all_feed", feed).Msg("all episodes feed is generated")
	feedMap["all"] = feed

	p.mu.Lock()
	p.feedMap = feedMap
	p.mu.Unlock()

	return nil
}

func (p *Podcaster) isAudioFile(fpath string) bool {
	f, err := os.Open(fpath)
	if err != nil {
		p.logger.Debug().Err(err).Str("path", fpath).
			Msg("failed to open file for checking audio file")
		return false
	}
	defer f.Close()

	// NOTE: 音声ファイルかどうかの判別には先頭20バイトあれば足りる
	head := make([]byte, 20)
	if _, err := f.Read(head); err != nil {
		p.logger.Debug().Err(err).Str("path", fpath).
			Msg("failed to read first 20 bytes of the file for checking audio file")
		return false
	}
	return filetype.IsAudio(head)
}
