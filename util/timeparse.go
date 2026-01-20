package util

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func parseTime(s string) (time.Time, error) {
	if strings.HasSuffix(s, "h") {
		d, _ := time.ParseDuration(s)
		return time.Now().Add(d), nil
	}
	return time.Parse(time.RFC3339, s)
}

func ParseFutureTime(s string) (time.Time, error) {
	now := time.Now()

	// 24h, 48h, etc.
	if strings.HasSuffix(s, "h") {
		d, err := time.ParseDuration(s)
		if err != nil {
			return time.Time{}, err
		}
		return now.Add(d), nil
	}

	// Unix timestamp
	if ts, err := strconv.ParseInt(s, 10, 64); err == nil {
		return time.Unix(ts, 0), nil
	}

	// RFC3339
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}

	return time.Time{}, errors.New("invalid time format")
}
