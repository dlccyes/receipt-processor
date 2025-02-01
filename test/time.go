package test

import "time"

func MustParseTime(timeStr string, format string) time.Time {
	t, err := time.Parse(format, timeStr)
	if err != nil {
		panic(err)
	}
	return t
}
