package config

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog"
)

const (
	AudioFormatAAC = "aac"
	AudioFormatMP3 = "mp3"
)

type Config struct {
	Programs []Program `yaml:"programs" json:"programs"`
}

type Program struct {
	Cron      string `yaml:"cron" json:"cron"`
	StationID string `yaml:"station" json:"station"`
	Start     string `yaml:"start" json:"start"`
	Encoding  string `yaml:"encoding,omitempty" json:"encoding,omitempty"`
}

func Parse(r io.Reader) (Config, error) {
	var c Config
	if err := yaml.NewDecoder(r).Decode(&c); err != nil {
		return c, errors.Wrap(err, "failed to decode yaml")
	}
	for i := range c.Programs {
		if c.Programs[i].Encoding == "" {
			c.Programs[i].Encoding = AudioFormatAAC
		}
	}
	return c, nil
}

func Init(configFilePath *string, configURL *string) (Config, error) {
	if configFilePath != nil {
		f, err := os.Open(*configFilePath)
		if err != nil {
			return Config{}, errors.Wrap(err, "failed to open config file")
		}
		defer f.Close()
		return Parse(f)
	}

	if configURL != nil {
		resp, err := http.Get(*configURL)
		if err != nil {
			return Config{}, errors.Wrap(
				err,
				fmt.Sprintf("failed to get config file: %s", *configURL),
			)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return Config{}, errors.Newf("status code is not 200: %d", resp.StatusCode)
		}
		return Parse(resp.Body)
	}

	return Config{}, nil
}

func (p Program) MarshalZerologObject(e *zerolog.Event) {
	e.Str("cron", p.Cron).
		Str("station_id", p.StationID).
		Str("start", p.Start).
		Str("encoding", p.Encoding)
}

type Programs []Program

func (p Programs) MarshalZerologArray(a *zerolog.Array) {
	for _, p := range p {
		a.Object(p)
	}
}

func (c Config) MarshalZerologObject(e *zerolog.Event) {
	e.Array("programs", Programs(c.Programs))
}
