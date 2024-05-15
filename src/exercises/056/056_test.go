package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestDigitSum(t *testing.T) {
	a := big.NewInt(1234)
	assert.Equal(t, 10, digitSum(a))
	a = big.NewInt(12345)
	assert.Equal(t, 15, digitSum(a))
}
