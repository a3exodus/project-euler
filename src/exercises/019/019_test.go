package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	// 9 Jan 1900 is a Tuesday
	allDays := compute(10)
	assert.Equal(t, 2, allDays[len(allDays)-1].dayOfWeek)

	// 9 Apr 1900 is a Monday
	allDays = compute(100)
	assert.Equal(t, 1, allDays[len(allDays)-1].dayOfWeek)

	// 26 September 2019 is a Friday
	allDays = compute(1_000)
	assert.Equal(t, 5, allDays[len(allDays)-1].dayOfWeek)

	// 18 May 1927 is a Wednesday
	allDays = compute(10_000)
	assert.Equal(t, 3, allDays[len(allDays)-1].dayOfWeek)
}
