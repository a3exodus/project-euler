package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestNPowerSumMod(t *testing.T) {
	assert.Equal(t, big.NewInt(2), nPowerSumMod(3, 2))
	assert.Equal(t, big.NewInt(5), nPowerSumMod(5, 3))
	assert.Equal(t, big.NewInt(2), nPowerSumMod(7, 4))
	assert.Equal(t, big.NewInt(110), nPowerSumMod(11, 5))

	ans := nPowerSumMod(71059, 7037)
	assert.Equal(t, 1, ans.Cmp(big.NewInt(1_000_000_000)))
}
