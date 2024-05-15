package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(100)
	assert.Equal(t, int64(44), ans)
}
