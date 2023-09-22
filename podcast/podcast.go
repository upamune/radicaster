package podcast

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Podcast struct {
	Title       string
	Link        string
	Description string
	PublishedAt *time.Time

	Episodes []Episode
}

type Episode struct {
	Title         string
	Description   string
	PublishedAt   *time.Time
	URL           string
	LengthInBytes int64
}

type Podcaster struct {
	logger zerolog.Logger

	baseURL   string
	targetDir string

	title       string
	link        string
	description string
	publishedAt *time.Time

	mu   *sync.RWMutex
	feed string
}

func NewPodcaster(
	logger zerolog.Logger,
	baseURL string,
	targetDir string,
	title string,
	link string,
	description string,
	publishedAt *time.Time,
) *Podcaster {
	p := &Podcaster{
		logger:      logger,
		baseURL:     baseURL,
		targetDir:   targetDir,
		title:       title,
		link:        link,
		description: description,
		publishedAt: publishedAt,
		mu:          &sync.RWMutex{},
	}
	return p
}

func (p *Podcaster) GetFeed() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.feed
}

func (p *Podcaster) Sync() error {
	p.logger.Info().Msg("Podcaster.Sync started")
	defer func() {
		p.logger.Info().Msg("Podcaster.Sync ended")
	}()

	podcast := &Podcast{
		Title:       p.title,
		Link:        p.link,
		Description: p.description,
		PublishedAt: p.publishedAt,
	}

	p.logger.Info().Str("target_dir", p.targetDir).Msg("filepath.Walk is starting")
	if err := filepath.Walk(p.targetDir, func(fpath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		p.logger.Info().Str("path", fpath).Err(err).Msg("found a target file")

		stat, err := os.Stat(fpath)
		if err != nil {
			return err
		}

		baseName := filepath.Base(fpath)

		u, err := url.Parse(p.baseURL)
		if err != nil {
			return fmt.Errorf("failed to parse baseURL(%s): %w", p.baseURL, err)
		}
		u.Path = path.Join(u.Path, "static", baseName)

		ep := Episode{
			Title:         fpath,
			URL:           u.String(),
			LengthInBytes: stat.Size(),
		}
		if ss := strings.Split(baseName, "_"); len(ss) > 1 {
			ep.Title = ss[0]
			startedAt, err := time.Parse("200601021504", strings.TrimSuffix(ss[1], filepath.Ext(ss[1])))
			if err != nil {
				return err
			}
			ep.PublishedAt = &startedAt
		}
		podcast.Episodes = append(podcast.Episodes, ep)

		return nil
	}); err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	p.logger.Info().Msg("encodeXML is starting")
	if err := encodeXML(buf, podcast); err != nil {
		return fmt.Errorf("failed to encodeXML: %w", err)
	}
	feed := buf.String()

	p.mu.Lock()
	p.feed = feed
	p.mu.Unlock()

	return nil
}
