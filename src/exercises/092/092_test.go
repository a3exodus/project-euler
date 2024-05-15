package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumDigitsSquared(t *testing.T) {
	assert.Equal(t, 32, sumDigitsSquared(44))
	assert.Equal(t, 13, sumDigitsSquared(32))
	assert.Equal(t, 10, sumDigitsSquared(13))
	assert.Equal(t, 1, sumDigitsSquared(1))
}

func TestArrivesAt89(t *testing.T) {
	assert.Equal(t, false, arrivesAt89(44))
	assert.Equal(t, true, arrivesAt89(85))
}
