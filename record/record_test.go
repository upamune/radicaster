package record

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
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
		c,
		"",
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
	}); err != nil {
		t.Fatalf("%+v\n", errors.WithStack(err))
	}
}
