package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFractionValue(t *testing.T) {
	assert.Equal(t, float64(1)/2, FractionValue(Fraction{1, 2}))
	assert.Equal(t, float64(1)/3, FractionValue(Fraction{1, 3}))
	assert.Equal(t, float64(1)/6, FractionValue(Fraction{1, 6}))
	assert.Equal(t, float64(1)/7, FractionValue(Fraction{1, 7}))
}
