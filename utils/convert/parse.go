package convert

import (
	"strconv"
	"time"
)

// ParseInteger is a wrapper function to return an int
// by parsing a given string. strconv.ParseInt() returns int64
// In case of an error, it returns the value x
func ParseInteger(s string, x int64) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return x
	}
	return int64(n)
}

// ParseFloat is a wrapper function to return an float
// by parsing a given string. strconv.ParseFloat() returns floatt64
// In case of an error, it returns the value x
func ParseFloat(s string, x float64) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return x
	}
	return float64(n)
}

// ParseTime parses the given value using the layout provided
// and time.Parse() utility
// If there is an Error, Return a Default Date = Jan 01 1970
func ParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		return time.Date(1970, time.January, 1, 2, 3, 4, 5, time.UTC)
	}
	return t
}
