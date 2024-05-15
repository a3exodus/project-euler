package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPerimeters(t *testing.T) {
	perimeters := getPerimeters(1000)

	assert.Equal(t, 1, perimeters[12])
	assert.Equal(t, 1, perimeters[24])
	assert.Equal(t, 1, perimeters[30])
	assert.Equal(t, 1, perimeters[36])
	assert.Equal(t, 1, perimeters[40])
	assert.Equal(t, 1, perimeters[48])

	assert.Equal(t, 0, perimeters[20])

	assert.Equal(t, 3, perimeters[120])
}
