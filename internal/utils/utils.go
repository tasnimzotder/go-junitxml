package utils

import (
	"strconv"
	"time"
)

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

func Float64ToSting(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
