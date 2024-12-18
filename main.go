package timephrase

import (
	"time"
)

func Parse(dateStr string, base time.Time) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, ErrEmptyString
	}

	t, err := parseAbsolute(dateStr)
	if err == nil {
		return t, nil
	}

	t, err = parseRelative(dateStr, base)
	if err == nil {
		return t, nil
	}

	return time.Time{}, err
}
