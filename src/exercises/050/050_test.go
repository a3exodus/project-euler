package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	ans := subSliceSum(slice, 1, 4)
	assert.Equal(t, 9, ans)
}
