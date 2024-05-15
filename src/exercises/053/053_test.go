package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestChoose(t *testing.T) {
	ans := choose(3, 1)
	expected := big.NewInt(3)
	assert.Equal(t, 0, ans.Cmp(expected))

	ans = choose(5, 2)
	expected = big.NewInt(10)
	assert.Equal(t, 0, ans.Cmp(expected))
}
