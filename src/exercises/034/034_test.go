package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := digitFactorialSum(145)
	assert.Equal(t, 145, ans)
}
