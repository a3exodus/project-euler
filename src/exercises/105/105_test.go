package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSubsetSum(t *testing.T) {
	A := []int{3, 5, 6, 7}
	assert.Equal(t, subsetSum{1, 3}, getSubsetSum(A, "1000"))
	assert.Equal(t, subsetSum{2, 8}, getSubsetSum(A, "1100"))
	assert.Equal(t, subsetSum{3, 14}, getSubsetSum(A, "1110"))
	assert.Equal(t, subsetSum{4, 21}, getSubsetSum(A, "1111"))
}

func TestIsSpecialSumSet(t *testing.T) {
	assert.Equal(t, true, isSpecialSumSet([]int{1}))
	assert.Equal(t, true, isSpecialSumSet([]int{1, 2}))
	assert.Equal(t, true, isSpecialSumSet([]int{2, 3, 4}))
	assert.Equal(t, true, isSpecialSumSet([]int{3, 5, 6, 7}))
	assert.Equal(t, true, isSpecialSumSet([]int{6, 9, 11, 12, 13}))

	assert.Equal(t, false, isSpecialSumSet([]int{1, 1}))
	assert.Equal(t, false, isSpecialSumSet([]int{2, 3, 2}))
	assert.Equal(t, false, isSpecialSumSet([]int{2, 3, 4, 5}))
}
