package utils

import (
	"strconv"
	"time"
)

// StringToDuration converts a string representation of time duration to a time.Duration value.
// If the timeString is empty, it returns 0.
// If the timeString does not have a "s" suffix, it adds the suffix to convert it to seconds (e.g. "1.2s").
// If the timeString is not a valid duration, it returns 0.
func StringToDuration(timeString string) time.Duration {
	if timeString == "" {
		return 0
	}

	// add "s" suffix to the string to convert it to seconds (e.g. "1.2s") if not empty string
	if timeString[len(timeString)-1] != 's' {
		timeString += "s"
	}

	duration, err := time.ParseDuration(timeString)
	if err != nil {
		return 0
	}

	return duration
}

// Float64ToSting converts a float64 value to a string representation.
// It uses strconv.FormatFloat to convert the float64 value to a string.
// The 'f' format specifier is used to format the float with no exponent and as many digits as necessary.
// The resulting string is returned.
func Float64ToSting(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
