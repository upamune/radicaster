package config

import (
	"io"

	"github.com/cockroachdb/errors"
	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog"
)

const (
	AudioFormatAAC = "aac"
	AudioFormatMP3 = "mp3"
)

type Config struct {
	Programs []Program `yaml:"programs"`
}

type Program struct {
	Cron      string `yaml:"cron"`
	StationID string `yaml:"station_id"`
	Start     string `yaml:"start"`
	Encoding  string `yaml:"encoding,omitempty"`
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
