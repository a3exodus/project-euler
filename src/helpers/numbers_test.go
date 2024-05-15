package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, IsPalindrome("12321"))
	assert.Equal(t, true, IsPalindrome("123321"))
	assert.Equal(t, false, IsPalindrome("12"))
	assert.Equal(t, false, IsPalindrome("123"))
}

func TestFactorize(t *testing.T) {
	assert.Equal(t, map[int64]int{2: 4}, Factorize(16))
	assert.Equal(t, map[int64]int{2: 1, 3: 1, 5: 1}, Factorize(30))
	assert.Equal(t, map[int64]int{2: 2, 5: 2}, Factorize(100))
	assert.Equal(t, map[int64]int{2: 2, 5: 2}, Factorize(100))
}

func TestNumDivisors(t *testing.T) {
	assert.Equal(t, int64(2), NumDivisors(3))
	assert.Equal(t, int64(5), NumDivisors(16))
	assert.Equal(t, int64(8), NumDivisors(30))
	assert.Equal(t, int64(9), NumDivisors(100))
}

func TestAreDigitPermutations(t *testing.T) {
	assert.Equal(t, true, AreDigitPermutations(int64(12345), int64(54321)))
	assert.Equal(t, true, AreDigitPermutations(int64(11122), int64(12121)))
	assert.Equal(t, false, AreDigitPermutations(int64(11), int64(22)))
	assert.Equal(t, false, AreDigitPermutations(int64(123), int64(124)))
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, 6, Factorial(3))
	assert.Equal(t, 3628800, Factorial(10))
	assert.Equal(t, 39916800, Factorial(11))
}

func TestSumDivisors(t *testing.T) {
	assert.Equal(t, 28, SumDivisors(28))
	assert.Equal(t, 220, SumDivisors(284))
	assert.Equal(t, 284, SumDivisors(220))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, int64(12), Lcm(4, 6))
	assert.Equal(t, int64(42), Lcm(21, 6))
	assert.Equal(t, int64(720), Lcm(48, 180))
}

func TestGcd(t *testing.T) {
	assert.Equal(t, int64(2), Gcd(4, 6))
	assert.Equal(t, int64(3), Gcd(21, 6))
	assert.Equal(t, int64(12), Gcd(48, 180))
}

func TestPhi(t *testing.T) {
	expected := map[int]int{
		2:  1,
		3:  2,
		4:  2,
		5:  4,
		6:  2,
		7:  6,
		8:  4,
		9:  6,
		10: 4,
	}

	for n, phiN := range expected {
		assert.Equal(t, phiN, Phi(n))
	}
}
