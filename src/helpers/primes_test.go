package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSieve(t *testing.T) {
	sieve := SieveOfEratosthenes(20)
	expected := []bool{
		false, false, true, true, false,
		true, false, true, false, false,
		false, true, false, true, false,
		false, false, true, false, true, false}
	assert.Equal(t, expected, sieve)
}
