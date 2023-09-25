package timeutil

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/goccy/go-yaml"
)

var (
	_ json.Marshaler        = (*Weekday)(nil)
	_ json.Unmarshaler      = (*Weekday)(nil)
	_ yaml.BytesMarshaler   = (*Weekday)(nil)
	_ yaml.BytesUnmarshaler = (*Weekday)(nil)

	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func JST() *time.Location {
	return jst
}

type Weekday time.Weekday

func (w *Weekday) String() string {
	if w == nil {
		return ""
	}
	return time.Weekday(*w).String()
}

func (w *Weekday) MarshalJSON() ([]byte, error) {
	if w == nil {
		return json.Marshal("")
	}
	return json.Marshal(time.Weekday(*w).String())
}

func (w *Weekday) UnmarshalJSON(b []byte) error {
	wd, err := newWeekday(string(b))
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal weekday")
	}
	*w = Weekday(wd)
	return nil
}

func (w *Weekday) MarshalYAML() ([]byte, error) {
	if w == nil {
		return yaml.Marshal("")
	}
	return yaml.Marshal(time.Weekday(*w).String())
}

func (w *Weekday) UnmarshalYAML(b []byte) error {
	wd, err := newWeekday(string(b))
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal weekday")
	}
	*w = Weekday(wd)
	return nil
}

func newWeekday(day string) (time.Weekday, error) {
	switch strings.ToLower(day) {
	case "sunday", "sun":
		return time.Sunday, nil
	case "monday", "mon":
		return time.Monday, nil
	case "tuesday", "tue":
		return time.Tuesday, nil
	case "wednesday", "wed":
		return time.Wednesday, nil
	case "thursday", "thu":
		return time.Thursday, nil
	case "friday", "fri":
		return time.Friday, nil
	case "saturday", "sat":
		return time.Saturday, nil
	default:
		return -1, fmt.Errorf("invalid day: %s", day)
	}
}

func LastSpecifiedWeekday(weekday Weekday, now time.Time) (time.Time, error) {
	targetWeekday := time.Weekday(weekday)
	for i := 0; i <= 7; i++ {
		previousDay := now.AddDate(0, 0, -i)
		if previousDay.Weekday() == targetWeekday {
			return previousDay, nil
		}
	}
	return time.Time{}, fmt.Errorf("failed to find last specified weekday: %s", time.Weekday(weekday))
}
