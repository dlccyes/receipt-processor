package test

import "time"

func MustParseDate(dateStr string) time.Time {
	return mustParseDateTime(dateStr, "2006-01-02")
}

func MustParseTime(timeStr string) time.Time {
	return mustParseDateTime(timeStr, "15:04")
}

func mustParseDateTime(timeStr string, format string) time.Time {
	t, err := time.Parse(format, timeStr)
	if err != nil {
		panic(err)
	}
	return t
}
