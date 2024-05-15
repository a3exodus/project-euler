package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBouncy(t *testing.T) {
	assert.Equal(t, false, isBouncy(134468))
	assert.Equal(t, false, isBouncy(66420))
	assert.Equal(t, true, isBouncy(155349))

	count := 0
	for i := 1; i <= 1_000; i++ {
		if isBouncy(i) {
			count++
		}
	}
	assert.Equal(t, 525, count)
}

func TestNumberBouncyNumbersBelow(t *testing.T) {
	assert.Equal(t, 538, getNForRatio(0.5))
	assert.Equal(t, 21780, getNForRatio(0.9))
}
