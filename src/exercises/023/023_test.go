package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 284, d(220))
	assert.Equal(t, 220, d(284))
}
