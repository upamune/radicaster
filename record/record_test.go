package record

import (
	"context"
	"fmt"
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

	startTime := time.Now().AddDate(0, 0, -1).Format("20060102")
	if err := r.Record(config.Program{
		Cron:      "",
		StationID: "LFR",
		Start:     fmt.Sprintf("%s0300", startTime),
		Encoding:  config.AudioFormatAAC,
	}); err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}
	time.Sleep(100 * time.Minute)
}
