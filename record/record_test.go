package record

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/upamune/podcast-server/config"
	"github.com/upamune/podcast-server/radikoutil"
)

func TestNewRecord(t *testing.T) {
	ctx := context.Background()
	c, err := radikoutil.NewClient(ctx)
	if err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}

	r, err := NewRecorder(
		zerolog.New(zerolog.NewConsoleWriter()),
		c,
		t.TempDir(),
		config.Config{
			Programs: []config.Program{},
		})
	if err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}

	if err := r.Record(config.Program{
		Cron:      "",
		StationID: "LFR",
		Start:     "0300",
		Encoding:  config.AudioFormatAAC,
	}); err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}
	time.Sleep(100 * time.Minute)
}
