package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(2)
	assert.Equal(t, int64(6), ans)
}
