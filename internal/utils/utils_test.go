package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringToDuration(t *testing.T) {
	// test case 1: empty string
	result := StringToDuration("")
	assert.Equal(t, time.Duration(0), result)

	// test case 2: valid time string
	result = StringToDuration("1.5")
	expected := time.Duration(1500000000)
	assert.Equal(t, expected, result)

	// test case 3: valid time string with suffix
	result = StringToDuration("1.5s")
	expected = time.Duration(1500000000)
	assert.Equal(t, expected, result)

	// test case 4: invalid time string
	result = StringToDuration("invalid")
	assert.Equal(t, time.Duration(0), result)
}
