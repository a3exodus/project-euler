package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiophantineReciprocals(t *testing.T) {
	assert.Equal(t, 1, diophantineReciprocals(1))
	assert.Equal(t, 2, diophantineReciprocals(2))
	assert.Equal(t, 2, diophantineReciprocals(3))
	assert.Equal(t, 3, diophantineReciprocals(4))
	assert.Equal(t, 2, diophantineReciprocals(5))
	assert.Equal(t, 5, diophantineReciprocals(6))
}
