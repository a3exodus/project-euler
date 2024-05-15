package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(1_000)
	assert.Equal(t, int64(4164), ans)
}
