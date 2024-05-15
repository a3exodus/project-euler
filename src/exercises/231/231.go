package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
The binomial coefficient (10 3) = 120.
120 = 2^3 x 3 x 5 = 2 x 2 x 2 x 3 x 5, and 2 + 2 + 2 + 3 + 5 = 14.
So the sum of the terms in the prime factorization of (10 3) is 13.
Find the sum of the terms in the prime factorization of (20_000_000 15_000_000).
*/
func compute(n int, k int) int64 {
	// Problem solved easily once you find the identity pS(a * b) = pS(a) + ps(b),
	// where pS(n) is the prime sum of n.
	totalSum := int64(0)
	for i := n - k + 1; i <= n; i++ {
		totalSum += int64(primeSum(i))
	}
	for i := 2; i <= k; i++ {
		totalSum -= int64(primeSum(i))
	}
	return totalSum
}

func primeSum(i int) int {
	sum := 0
	factors := helpers.Factorize(int64(i))
	for prime, power := range factors {
		sum += int(prime) * power
	}
	return sum
}

func main() {
	fmt.Println(compute(20_000_000, 15_000_000))
}
