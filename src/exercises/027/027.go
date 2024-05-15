package main

import (
	"fmt"
	"project-euler/helpers"
)

var sieve []bool

/*
Euler discovered the remarkable quadratic formula: n^2+n+41.

It turns out that the formula will produce 40 primes for the consecutive integer values 0 <= n =< 39. However, when n=40,
40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.
The incredible formula n^2 — 79n + 1601 was discovered, which produces 80 primes for the consecutive values 0 < n < 79. The product of the coefficients, —79 and 1601, is —126479.
Considering quadratics of the form: n^2 + an + b, where |a| < 1000 and |b| <= 1000, where |n| is the modulus/absolute value of n, e.g. |11| = 11 and |—4| = 4
Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/
func main() {
	sieve = helpers.SieveOfEratosthenes(1_000_000)
	//str := streak(1, 41, sieve)
	//str := streak(-79, 1601, sieve)

	maxStreak := 1
	maxProduct := 1
	for b := 2; b <= 1000; b++ {
		for a := -b; a < 1000; a++ {
			str := streak(a, b)
			if maxStreak < str {
				maxStreak = str
				maxProduct = a * b
			}
		}
	}
	fmt.Println(maxProduct)
}

func streak(a int, b int) int {
	val := 0
	for n := 0; n <= 1000; n++ {
		val = function(n, a, b)
		if !isPrime(val) {
			return n
		}
	}
	return -1
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	return sieve[n]
}

func function(n int, a int, b int) int {
	return n*n + a*n + b
}
