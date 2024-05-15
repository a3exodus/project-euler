package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := numSols(120)
	assert.Equal(t, 3, ans)
}
