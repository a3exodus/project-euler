package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(6)
	assert.Equal(t, 13, ans)
}
