package timeutil

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestLastSpecifiedWeekday(t *testing.T) {
	t.Parallel()
	tests := []struct {
		weekday Weekday
		want    string
	}{
		{
			weekday: Weekday(time.Monday),
			want:    "2023-09-25",
		},
		{
			weekday: Weekday(time.Sunday),
			want:    "2023-09-24",
		},
		{
			weekday: Weekday(time.Saturday),
			want:    "2023-09-23",
		},
		{
			weekday: Weekday(time.Friday),
			want:    "2023-09-22",
		},
		{
			weekday: Weekday(time.Thursday),
			want:    "2023-09-21",
		},
		{
			weekday: Weekday(time.Wednesday),
			want:    "2023-09-20",
		},
		{
			weekday: Weekday(time.Tuesday),
			want:    "2023-09-19",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("last %s", time.Weekday(tt.weekday)), func(t *testing.T) {
			t.Parallel()
			now := time.Date(2023, time.September, 25, 0, 3, 0, 0, JST())
			got, err := LastSpecifiedWeekday(tt.weekday, now)
			if err != nil {
				t.Fatalf("LastSpecifiedWeekday() error = %v", err)
				return
			}
			if got, want := got.Format(time.DateOnly), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("LastSpecifiedWeekday() got = %v, want %v", got, want)
			}
		})
	}
}
