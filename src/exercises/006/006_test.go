package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(10)
	assert.Equal(t, int64(2640), ans)
}
