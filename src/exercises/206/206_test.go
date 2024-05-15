package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsConcealedSquare(t *testing.T) {
	ans := IsConcealedSquare(11223344556677889)
	assert.Equal(t, true, ans)
	ans = IsConcealedSquare(22334455667788990)
	assert.Equal(t, false, ans)
}
