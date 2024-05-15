package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRad(t *testing.T) {
	expected := [20]int{
		1, 2, 4, 8, 16, 3, 9, 5, 6, 12, 18, 7, 10, 20, 11, 13, 14, 15, 17, 19}
	fmt.Println(expected)
	n := 20
	ans := compute(n)
	for i := 1; i <= n; i++ {
		assert.Equal(t, expected[i-1], ans[i-1].index)
	}
}
