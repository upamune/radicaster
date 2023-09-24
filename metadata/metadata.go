package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cockroachdb/errors"
)

type EpisodeMetadata struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	ImageURL    string    `json:"image_url"`
}

func createMetadataPath(base string) string {
	return fmt.Sprintf("%s.json", base)
}

func WriteByAudioFilePath(basePath string, metadata EpisodeMetadata) error {
	path := createMetadataPath(basePath)

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "failed to create metadata file")
	}

	if err := json.NewEncoder(f).Encode(metadata); err != nil {
		return errors.Wrap(err, "failed to encode metadata json")
	}
	return nil
}

func ReadByAudioFilePath(basePath string) (EpisodeMetadata, error) {
	path := createMetadataPath(basePath)

	f, err := os.Open(path)
	if err != nil {
		return EpisodeMetadata{}, errors.Wrap(err, "failed to read metadata file")
	}
	defer f.Close()

	var meta EpisodeMetadata
	if json.NewDecoder(f).Decode(&meta); err != nil {
		return EpisodeMetadata{}, errors.Wrap(err, "failed to decode metadata json")
	}
	return meta, nil
}
