package record

import (
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"github.com/upamune/radicaster/config"
	"github.com/upamune/radicaster/timeutil"
)

func TestNewRecord(t *testing.T) {
	r, err := NewRecorder(
		zerolog.New(zerolog.NewConsoleWriter()).Level(zerolog.DebugLevel),
		t.TempDir(),
		"",
		"",
		config.Config{
			Programs: []config.Program{},
		},
		"",
	)
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

func TestRecorder_RecordAll(t *testing.T) {
	t.Parallel()
	r, err := NewRecorder(zerolog.Nop(), t.TempDir(), "", "", config.Config{}, "")
	if err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}
	if err := r.RecordAll(); err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}
}
