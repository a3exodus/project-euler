package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestIsLychrel(t *testing.T) {
	assert.Equal(t, true, isLychrel(big.NewInt(196)))   // is  Lychrel number.
	assert.Equal(t, false, isLychrel(big.NewInt(349)))  // requires 3 < 50 iterations.
	assert.Equal(t, true, isLychrel(big.NewInt(10677))) // requires 53 > 50 iterations.
}
