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
	"strconv"
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

	mu                  *sync.RWMutex
	feedMap             map[string]string
	pathGroupedEpisodes map[string][]Episode
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
	return p.GetDefaultFeedWithSince("")
}

func (p *Podcaster) GetDefaultFeedWithSince(since string) string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	if since == "" {
		return p.feedMap[""]
	}

	// Parse the since duration
	duration, err := parseDurationString(since)
	if err != nil {
		p.logger.Warn().Err(err).Str("since", since).Msg("failed to parse since duration, returning unfiltered feed")
		return p.feedMap[""]
	}

	// Generate filtered feed for default path
	feed, ok := p.generateFilteredFeed("", duration)
	if !ok {
		return p.feedMap[""]
	}
	return feed
}

func (p *Podcaster) GetFeed(path string) (string, bool) {
	return p.GetFeedWithSince(path, "")
}

func (p *Podcaster) GetFeedWithSince(path string, since string) (string, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	feed, ok := p.feedMap[path]
	if !ok {
		return "", false
	}

	// If no since filter, return the original feed
	if since == "" {
		return feed, true
	}

	// Parse the since duration
	duration, err := parseDurationString(since)
	if err != nil {
		p.logger.Warn().Err(err).Str("since", since).Msg("failed to parse since duration, returning unfiltered feed")
		return feed, true
	}

	// For filtered feeds, we need to regenerate the feed with filtered episodes
	return p.generateFilteredFeed(path, duration)
}

func (p *Podcaster) generateFilteredFeed(path string, since time.Duration) (string, bool) {
	var episodes []Episode
	var ok bool

	// Handle "all" episodes case
	if path == "all" {
		// Collect all episodes from all paths
		for _, pathEpisodes := range p.pathGroupedEpisodes {
			episodes = append(episodes, pathEpisodes...)
		}
		ok = len(episodes) > 0
	} else {
		// Get episodes from the specific path
		episodes, ok = p.getEpisodesForPath(path)
	}

	if !ok {
		return "", false
	}

	// Filter episodes by the since duration
	filteredEpisodes := filterEpisodesBySince(episodes, since)
	if len(filteredEpisodes) == 0 {
		// Return empty feed if no episodes match the filter
		return p.generateEmptyFeed(path)
	}

	// Sort filtered episodes
	sortEpisodesByPublishedAtDesc(filteredEpisodes)
	
	// Generate the podcast feed with filtered episodes
	return p.generatePodcastFeed(path, filteredEpisodes)
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

// parseDurationString parses duration strings like "1y", "6m", "30d", "24h"
func parseDurationString(durationStr string) (time.Duration, error) {
	if durationStr == "" {
		return 0, nil
	}

	durationStr = strings.ToLower(strings.TrimSpace(durationStr))
	if len(durationStr) < 2 {
		return 0, fmt.Errorf("invalid duration format: %s", durationStr)
	}

	numStr := durationStr[:len(durationStr)-1]
	unit := durationStr[len(durationStr)-1:]

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("invalid number in duration: %s", durationStr)
	}

	switch unit {
	case "h":
		return time.Duration(num) * time.Hour, nil
	case "d":
		return time.Duration(num) * 24 * time.Hour, nil
	case "m":
		return time.Duration(num) * 30 * 24 * time.Hour, nil // approximate month as 30 days
	case "y":
		return time.Duration(num) * 365 * 24 * time.Hour, nil // approximate year as 365 days
	default:
		return 0, fmt.Errorf("unsupported duration unit: %s", unit)
	}
}

// filterEpisodesBySince filters episodes published since the given duration ago
func filterEpisodesBySince(episodes []Episode, since time.Duration) []Episode {
	if since == 0 {
		return episodes
	}

	cutoff := time.Now().Add(-since)
	filtered := make([]Episode, 0, len(episodes))

	for _, ep := range episodes {
		if ep.PublishedAt != nil && ep.PublishedAt.After(cutoff) {
			filtered = append(filtered, ep)
		}
	}

	return filtered
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

		// targetDir からの相対パスを取得（サブディレクトリを含む）
		relPath, err := filepath.Rel(p.targetDir, fpath)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}
		baseName := filepath.Base(fpath)

		u, err := url.Parse(p.baseURL)
		if err != nil {
			return fmt.Errorf("failed to parse baseURL(%s): %w", p.baseURL, err)
		}
		// path.Join を使ってパスを結合（forward slash を使用）
		u.Path = path.Join(u.Path, "static", filepath.ToSlash(relPath))

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

	// adhoc パスの空のフィードを初期化（録音ファイルがなくてもRSSフィードが存在するように）
	if _, exists := feedMap["adhoc"]; !exists {
		adhocEpisodes := pathGroupedEpisodes["adhoc"]
		if adhocEpisodes == nil {
			adhocEpisodes = []Episode{}
		}
		adhocFeed, err := encodePodcastToXML(
			&Podcast{
				Title:       "Radicaster - アドホック録音",
				Link:        p.link,
				Description: "番組表から手動で録音した番組",
				PublishedAt: p.publishedAt,
				ImageURL:    p.imageURL,
				Episodes:    adhocEpisodes,
			},
		)
		if err == nil {
			feedMap["adhoc"] = adhocFeed
		}
	}

	p.mu.Lock()
	p.feedMap = feedMap
	p.pathGroupedEpisodes = pathGroupedEpisodes
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

func (p *Podcaster) getEpisodesForPath(path string) ([]Episode, bool) {
	episodes, ok := p.pathGroupedEpisodes[path]
	return episodes, ok
}

func (p *Podcaster) generateEmptyFeed(path string) (string, bool) {
	podcast := p.createPodcastForPath(path, []Episode{})
	buf := bytes.NewBuffer(nil)
	if err := encodeXML(buf, podcast); err != nil {
		p.logger.Err(err).Str("path", path).Msg("failed to encode empty feed")
		return "", false
	}
	return buf.String(), true
}

func (p *Podcaster) generatePodcastFeed(path string, episodes []Episode) (string, bool) {
	podcast := p.createPodcastForPath(path, episodes)
	buf := bytes.NewBuffer(nil)
	if err := encodeXML(buf, podcast); err != nil {
		p.logger.Err(err).Str("path", path).Msg("failed to encode podcast feed")
		return "", false
	}
	return buf.String(), true
}

func (p *Podcaster) createPodcastForPath(path string, episodes []Episode) *Podcast {
	podcast := &Podcast{
		Title:       p.title,
		Link:        p.link,
		Description: p.description,
		PublishedAt: p.publishedAt,
		ImageURL:    p.imageURL,
		Episodes:    episodes,
	}

	// For specific paths, customize the podcast based on the first episode
	if path != "" && len(episodes) > 0 {
		latestEpisode := episodes[0]
		
		podcastTitle := p.title
		if latestEpisode.PodcastTitle != "" {
			podcastTitle = latestEpisode.PodcastTitle
		}
		
		podcast = &Podcast{
			Title:       podcastTitle,
			Link:        p.link,
			Description: p.description,
			PublishedAt: p.publishedAt,
			ImageURL:    latestEpisode.ImageURL,
			Episodes:    episodes,
		}
	}

	// Special handling for "all" path
	if path == "all" {
		podcast.Title = fmt.Sprintf("%s(ALL)", p.title)
		podcast.ImageURL = p.imageURL
	}

	return podcast
}
