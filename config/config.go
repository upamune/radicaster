package config

import (
	"io"

	"github.com/cockroachdb/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Programs []Program `yaml:"programs"`
}

type Program struct {
	Cron      string `yaml:"cron"`
	StationID string `yaml:"station_id"`
	Start     string `yaml:"start"`
}

func Parse(r io.Reader) (Config, error) {
	var c Config
	if err := yaml.NewDecoder(r).Decode(&c); err != nil {
		return c, errors.Wrap(err, "failed to decode yaml")
	}
	return c, nil
}
