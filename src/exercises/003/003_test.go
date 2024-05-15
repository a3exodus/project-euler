package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(13195)
	assert.Equal(t, int64(29), ans)
}
