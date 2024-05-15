package helpers

import (
	"math"
	"strconv"
)

/*
input: s string
output: true if and only if s is palindrome
*/
func IsPalindrome(s string) bool {
	var reverse string
	for _, v := range s {
		reverse = string(v) + reverse
	}
	return s == reverse
}

/*
input: n int64
output: slice of the factors of n
*/
func Factorize(n int64) map[int64]int {
	factors := map[int64]int{}
	for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors[i]++
			n /= i
			i--
		}
	}
	factors[n]++
	return factors
}

/*
input: n int64
output: the number of divisor of n
*/
func NumDivisors(n int64) int64 {
	factors := Factorize(int64(n))
	numDivisors := int64(1)
	for _, power := range factors {
		numDivisors *= int64(power + 1)
	}
	return numDivisors
}

/*
input: a int64, b int64
output: true iff a and b are permutations of one another
*/
func AreDigitPermutations(a int64, b int64) bool {
	digitCount := [10]int{}
	numA := strconv.FormatInt(a, 10)
	numB := strconv.FormatInt(b, 10)
	n := len(numA)
	if n != len(numB) {
		return false
	}

	for i := 0; i < n; i++ {
		digitA := int(numA[i] - '0')
		digitCount[digitA]++
		digitB := int(numB[i] - '0')
		digitCount[digitB]--
	}

	for i := 0; i < 10; i++ {
		if digitCount[i] != 0 {
			return false
		}
	}
	return true
}

/*
input: n int
output: n!
*/
func Factorial(n int) int {
	if n < 10 {
		factorials := [10]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
		return factorials[n]
	}
	return n * Factorial(n-1)
}

/*
input: n int
output: sum divisors of n
*/
func SumDivisors(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

/*
input: a int64, b int64
output: least common multiple of a and b
*/
func Lcm(a int64, b int64) int64 {
	return a * b / Gcd(a, b)
}

/*
input: a int64, b int64
output: greatest common divisor of a and b
*/
func Gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Phi(n int) int {
	factors := Factorize(int64(n))
	phiN := 1
	for p, k := range factors {
		phiN *= int(math.Pow(float64(p), float64(k-1))) * int(p-1)
	}
	return phiN
}
