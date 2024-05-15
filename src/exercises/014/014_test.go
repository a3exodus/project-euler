package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ans := getCollatzChainLength(13)
	assert.Equal(t, 10, ans)
}
