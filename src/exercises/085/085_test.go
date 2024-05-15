package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberRectangles(t *testing.T) {
	ans := getNumberRectangles(2, 3)
	assert.Equal(t, int64(18), ans)
}
