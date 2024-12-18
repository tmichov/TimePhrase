package timephrase

import (
	"testing"
	"time"

	timephrase "github.com/tmichov/TimePhrase"
)

func TestParseAbsolute(t *testing.T) {
	base := time.Now()
	expected := time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)

	result, err := timephrase.Parse("2024-12-25", base)
	if err != nil {
		t.Fatalf("Error parsing time: %v", err)
	}

	if !result.Equal(expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestParseRelative(t *testing.T) {
	base := time.Date(2024, 12, 18, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		input    string
		expected time.Time
	}{
		{"today", base},
		{"yesterday", base.AddDate(0, 0, -1)},
		{"tomorrow", base.AddDate(0, 0, 1)},
		{"2 days ago", base.AddDate(0, 0, -2)},
		{"5 hours ago", base.Add(-5 * time.Hour)},
		{"3 weeks from now", base.AddDate(0, 0, 3*7)},
		{"3 weeks ago", base.AddDate(0, 0, -3*7)},
	}

	for _, test := range tests {
		result, err := timephrase.Parse(test.input, base)
		if err != nil {
			t.Fatalf("Failed to parse '%s': %v", test.input, err)
		}
		if !result.Equal(test.expected) {
			t.Errorf("For '%s', expected %v, got %v", test.input, test.expected, result)
		}
	}
}
