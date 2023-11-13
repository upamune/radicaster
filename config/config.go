package config

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog"
	"github.com/upamune/radicaster/timeutil"
)

const (
	AudioFormatAAC = "aac"
	AudioFormatMP3 = "mp3"

	zenrokuDefaultCronExpression = "0 3 * * *"
)

type Config struct {
	Programs []Program `yaml:"programs" json:"programs"`
	Zenroku  Zenroku   `yaml:"zenroku" json:"zenroku"`
}

type Stations map[string]Station

type Zenroku struct {
	Enable           bool     `yaml:"enable" json:"enable"`
	Cron             string   `yaml:"cron" json:"cron"`
	Encoding         string   `yaml:"encoding" json:"encoding"`
	Stations         Stations `yaml:"stations" json:"stations"`
	EnableStationIDs []string `yaml:"enable_stations" json:"enable_stations"`
}

type Station struct {
	ImageURL string `yaml:"image_url" json:"image_url"`
}

type Program struct {
	Title     string             `yaml:"title" json:"title"`
	Weekdays  []timeutil.Weekday `yaml:"weekdays" json:"weekdays"`
	Cron      string             `yaml:"cron" json:"cron"`
	AreaID    string             `yaml:"area_id,omitempty" json:"area_id,omitempty"`
	StationID string             `yaml:"station" json:"station"`
	Start     string             `yaml:"start" json:"start"`
	Encoding  string             `yaml:"encoding" json:"encoding"`
	ImageURL  string             `yaml:"image_url" json:"image_url"`
	Path      string             `yaml:"path" json:"path"`
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
	if c.Zenroku.Cron == "" {
		c.Zenroku.Cron = zenrokuDefaultCronExpression
	}
	if c.Zenroku.Encoding == "" {
		c.Zenroku.Encoding = AudioFormatAAC
	}
	for k, v := range c.Zenroku.Stations {
		k, v := k, v
		// NOTE: 小文字でも引けるようにする
		c.Zenroku.Stations[strings.ToLower(k)] = v
	}
	return c, nil
}

func Init(configFilePath *string, configURL *string) (Config, error) {
	if configFilePath != nil && *configFilePath != "" {
		// NOTE: ファイルが存在しない場合、空で作成する
		if _, err := os.Stat(*configFilePath); err != nil {
			f, err := os.Create(*configFilePath)
			if err != nil {
				return Config{}, errors.Wrap(err, "failed to create config file")
			}
			defer f.Close()
			var c Config
			if err := yaml.NewEncoder(f).Encode(c); err != nil {
				return c, errors.Wrap(err, "failed to encode yaml")
			}
			return c, nil
		}
		f, err := os.Open(*configFilePath)
		if err != nil {
			return Config{}, errors.Wrap(err, "failed to open config file")
		}
		defer f.Close()
		return Parse(f)
	}

	if configURL != nil && *configURL != "" {
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

func (s Station) MarshalZerologObject(e *zerolog.Event) {
	e.Str("image_url", s.ImageURL)
}

func (s Stations) MarshalZerologObject(e *zerolog.Event) {
	for sid, s := range s {
		e.Object(sid, s)
	}
}

func (z Zenroku) MarshalZerologObject(e *zerolog.Event) {
	e.Str("cron", z.Cron).
		Str("encoding", z.Encoding).
		Bool("enable", z.Enable).
		Strs("enable_station_ids", z.EnableStationIDs).
		Object("stations", z.Stations)
}

func (p Program) MarshalZerologObject(e *zerolog.Event) {
	e.Str("cron", p.Cron).
		Str("title", p.Title).
		Str("station_id", p.StationID).
		Str("start", p.Start).
		Str("encoding", p.Encoding).
		Str("image_url", p.ImageURL).
		Str("path", p.Path)
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

func (c Config) Validate() error {
	for _, program := range c.Programs {
		p := strings.ToLower(strings.TrimPrefix(program.Path, "/"))
		if p == "all" {
			return errors.Newf(
				"pathに `all` は使用できません: program_title=%s",
				program.Title,
			)
		}
	}
	return nil
}
