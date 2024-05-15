package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(10)
	expected := big.NewInt(int64(10405071317))
	expected.Mod(expected, big.NewInt(10_000_000_000))
	assert.Equal(t, 0, ans.Cmp(expected))
}
