package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSeq(t *testing.T) {
	assert.Equal(t, "192384576", generateSeq(192))
	assert.Equal(t, "918273645", generateSeq(9))
}
