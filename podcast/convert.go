package podcast

import (
	"fmt"
	"io"
	"path"
	"time"

	podcastpkg "github.com/eduncan911/podcast"
)

func encodeXML(w io.Writer, p *Podcast) error {
	now := time.Now()
	podcast := podcastpkg.New(
		p.Title,
		p.Link,
		p.Description,
		p.PublishedAt,
		&now,
	)
	for _, e := range p.Episodes {
		e := e
		desc := e.Description
		if desc == "" {
			desc = "dummy description"
		}
		item := podcastpkg.Item{
			Title:       e.Title,
			Description: desc,
			PubDate:     e.PublishedAt,
		}

		enclosureType := podcastpkg.M4A
		if path.Ext(e.URL) == ".mp3" {
			enclosureType = podcastpkg.MP3
		}

		item.AddEnclosure(
			e.URL,
			enclosureType,
			e.LengthInBytes,
		)

		if _, err := podcast.AddItem(item); err != nil {
			return fmt.Errorf("failed to add item to podcast: %w", err)
		}
	}

	return podcast.Encode(w)
}
