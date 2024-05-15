package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberNonBouncyBelowPower10(t *testing.T) {
	assert.Equal(t, int64(9), numberNonBouncyBelowPower10(1))
	assert.Equal(t, int64(99), numberNonBouncyBelowPower10(2))

	// test case from exercise 112
	assert.Equal(t, int64(474), numberNonBouncyBelowPower10(3))

	assert.Equal(t, int64(12951), numberNonBouncyBelowPower10(6))
	assert.Equal(t, int64(277032), numberNonBouncyBelowPower10(10))
}
