package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestRMax(t *testing.T) {
	rMaximum := rMax(7)
	expected := big.NewInt(42)
	assert.Equal(t, 0, expected.Cmp(rMaximum))
}
