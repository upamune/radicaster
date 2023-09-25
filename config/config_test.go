package config

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/upamune/podcast-server/timeutil"
)

func TestParse(t *testing.T) {
	t.Parallel()
	wantConfig := Config{
		Programs: []Program{
			{
				Weekdays: []timeutil.Weekday{
					timeutil.Weekday(time.Sunday),
					timeutil.Weekday(time.Monday),
					timeutil.Weekday(time.Tuesday),
				},
				Cron:      "10 3 * * 6",
				StationID: "LFR",
				Start:     "0100",
				Encoding:  "aac",
				ImageURL:  "http://example.com/image.png",
			},
			{
				Weekdays: []timeutil.Weekday{
					timeutil.Weekday(time.Friday),
				},
				Cron:      "40 4 * * 2",
				StationID: "LFR",
				Start:     "0300",
				Encoding:  "mp3",
			},
		},
	}
	tests := map[string]struct {
		filename string
	}{
		"yaml": {
			filename: "testdata/radicast.yaml",
		},
	}
	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r, err := os.Open(tt.filename)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
			}
			got, err := Parse(r)
			if err != nil {
				t.Fatalf("failed to parse: %v", err)
				return
			}
			if !reflect.DeepEqual(got, wantConfig) {
				t.Errorf("Parse() got = %v, want %v", got, wantConfig)
			}
		})
	}
}
