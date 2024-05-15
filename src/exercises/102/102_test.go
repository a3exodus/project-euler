package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointInTriangle(t *testing.T) {
	origin := vector{0, 0}
	assert.Equal(t, true, pointInTriangle(origin, vector{-340, 495}, vector{-153, -910}, vector{835, -947}))
	assert.Equal(t, false, pointInTriangle(origin, vector{-175, 41}, vector{-421, -714}, vector{574, -645}))
}
