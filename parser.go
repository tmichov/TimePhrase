package timephrase

import (
	"regexp"
	"strconv"
	"time"
)

func parseAbsolute(s string) (time.Time, error) {
	layouts := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02T15",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04Z",
		"2006-01-02T15Z",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04-07:00",
		"2006-01-02T15-07:00",
		"2006-01-02T15:04:05-0700",
		"2006-01-02T15:04-0700",
		"2006-01-02T15-0700",
		"2006-01-02T15:04:05-07",
		"2006-01-02T15:04-07",
		"2006-01-02T15-07",
		"02 Jan 2006",
		"02 Jan 2006 15:04:05",
		"02 Jan 2006 15:04",
		"02 Jan 2006 15",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, ErrAbsoluteParse
}

func parseRelative(s string, base time.Time) (time.Time, error) {
	if s == "now" {
		return base, nil
	} else if s == "today" {
		return base, nil
	} else if s == "tomorrow" {
		return base.AddDate(0, 0, 1), nil
	} else if s == "yesterday" {
		return base.AddDate(0, 0, -1), nil
	}

	reOffset := regexp.MustCompile(`([+-]\d+)\s*(day|week|month|year|hour|minute|second)s?`)
	if matches := reOffset.FindStringSubmatch(s); len(matches) == 3 {
		value, _ := strconv.Atoi(matches[1])
		unit := matches[2]
		return adjustDate(base, value, unit)
	}

	reAgo := regexp.MustCompile(`(\d+)\s*(day|week|month|year|hour|minute|second)s?\s+ago`)
	if matches := reAgo.FindStringSubmatch(s); len(matches) == 3 {
		value, _ := strconv.Atoi(matches[1])
		unit := matches[2]
		return adjustDate(base, -value, unit)
	}

	reFromNow := regexp.MustCompile(`(\d+)\s*(day|week|month|year|hour|minute|second)s?\s+from\s+now`)
	if matches := reFromNow.FindStringSubmatch(s); len(matches) == 3 {
		value, _ := strconv.Atoi(matches[1])
		unit := matches[2]
		return adjustDate(base, value, unit)
	}

	reIn := regexp.MustCompile(`in\s+(\d+)\s*(day|week|month|year|hour|minute|second)s?`)
	if matches := reIn.FindStringSubmatch(s); len(matches) == 3 {
		value, _ := strconv.Atoi(matches[1])
		unit := matches[2]
		return adjustDate(base, value, unit)
	}

	return time.Time{}, ErrRelativeParse
}

func adjustDate(base time.Time, value int, unit string) (time.Time, error) {
	switch unit {
	case "day":
		return base.AddDate(0, 0, value), nil
	case "week":
		return base.AddDate(0, 0, value*7), nil
	case "month":
		return base.AddDate(0, value, 0), nil
	case "year":
		return base.AddDate(value, 0, 0), nil
	case "hour":
		return base.Add(time.Duration(value) * time.Hour), nil
	case "minute":
		return base.Add(time.Duration(value) * time.Minute), nil
	case "second":
		return base.Add(time.Duration(value) * time.Second), nil
	default:
		return time.Time{}, ErrRelativeParse
	}
}
