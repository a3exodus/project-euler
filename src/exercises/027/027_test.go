package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunction(t *testing.T) {
	ans := function(40, 1, 41)
	assert.Equal(t, 41*41, ans)

	ans = function(41, 1, 41)
	assert.Equal(t, 41*43, ans)
}
