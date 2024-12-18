package timephrase

import (
	"time"
)

func Parse(dateStr string, base time.Time) (time.Time, error) {

	if dateStr == "" {
		return time.Time{}, ErrEmptyString
	}

	if t, err := parseAbsolute(dateStr); err == nil {
		return t, nil
	}

	if t, err := parseRelative(dateStr, base); err == nil {
		return t, nil
	}

	return time.Time{}, ErrAbsoluteParse
}
