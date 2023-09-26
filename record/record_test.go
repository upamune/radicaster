package record

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/upamune/radicaster/config"
	"github.com/upamune/radicaster/radikoutil"
	"github.com/upamune/radicaster/timeutil"
)

func TestNewRecord(t *testing.T) {
	ctx := context.Background()
	c, err := radikoutil.NewClient(ctx)
	if err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}

	r, err := NewRecorder(
		zerolog.New(zerolog.NewConsoleWriter()).Level(zerolog.DebugLevel),
		c,
		t.TempDir(),
		config.Config{
			Programs: []config.Program{},
		})
	if err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}

	now := time.Now().AddDate(0, 0, -1)
	if err := r.Record(config.Program{
		Cron:      "",
		Weekdays:  []timeutil.Weekday{timeutil.Weekday(now.Weekday())},
		StationID: "LFR",
		Start:     "0300",
		Encoding:  config.AudioFormatAAC,
	}); err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}
}
