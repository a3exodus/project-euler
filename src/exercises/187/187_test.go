package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := compute(30)
	assert.Equal(t, 10, ans)
}
