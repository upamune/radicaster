package podcast

import (
	"testing"
	"time"
)

func TestParseDurationString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		wantErr  bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "1 hour",
			input:    "1h",
			expected: time.Hour,
			wantErr:  false,
		},
		{
			name:     "24 hours",
			input:    "24h",
			expected: 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:     "1 day",
			input:    "1d",
			expected: 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:     "30 days",
			input:    "30d",
			expected: 30 * 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:     "1 month",
			input:    "1m",
			expected: 30 * 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:     "6 months",
			input:    "6m",
			expected: 6 * 30 * 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:     "1 year",
			input:    "1y",
			expected: 365 * 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:     "uppercase",
			input:    "1Y",
			expected: 365 * 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:     "with spaces",
			input:    " 1y ",
			expected: 365 * 24 * time.Hour,
			wantErr:  false,
		},
		{
			name:    "invalid format - too short",
			input:   "1",
			wantErr: true,
		},
		{
			name:    "invalid format - non-numeric",
			input:   "xy",
			wantErr: true,
		},
		{
			name:    "invalid unit",
			input:   "1x",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseDurationString(tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseDurationString() expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("parseDurationString() unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("parseDurationString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestFilterEpisodesBySince(t *testing.T) {
	now := time.Now()
	episodes := []Episode{
		{
			Title:       "Recent Episode",
			PublishedAt: &now,
		},
		{
			Title:       "Old Episode",
			PublishedAt: func() *time.Time { t := now.Add(-2 * 365 * 24 * time.Hour); return &t }(),
		},
		{
			Title:       "Medium Age Episode",
			PublishedAt: func() *time.Time { t := now.Add(-6 * 30 * 24 * time.Hour); return &t }(),
		},
		{
			Title:       "Episode with nil PublishedAt",
			PublishedAt: nil,
		},
	}

	tests := []struct {
		name     string
		since    time.Duration
		expected int
	}{
		{
			name:     "no filter",
			since:    0,
			expected: 4,
		},
		{
			name:     "1 year filter",
			since:    365 * 24 * time.Hour,
			expected: 2, // recent and medium age episodes
		},
		{
			name:     "1 hour filter",
			since:    time.Hour,
			expected: 1, // only recent episode
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filterEpisodesBySince(episodes, tt.since)
			if len(result) != tt.expected {
				t.Errorf("filterEpisodesBySince() returned %d episodes, expected %d", len(result), tt.expected)
			}
		})
	}
}