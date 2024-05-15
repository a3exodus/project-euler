package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAmicableChain(t *testing.T) {
	assert.Equal(t, 2, len(amicableChain(220)))
	assert.Equal(t, 5, len(amicableChain(12496)))
}
