package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorialChain(t *testing.T) {
	assert.Equal(t, 5, factorialChain(69))
	assert.Equal(t, 4, factorialChain(78))
	assert.Equal(t, 2, factorialChain(540))
	assert.Equal(t, 2, factorialChain(871))
	assert.Equal(t, 2, factorialChain(872))
}
