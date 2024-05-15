package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
func compute(n int) int64 {
	sum := int64(0)
	sieve := helpers.SieveOfEratosthenes(n)
	for i := 1; i < n; i++ {
		if sieve[i] {
			sum += int64(i)
		}
	}
	return sum
}

func main() {
	fmt.Println(compute(2_000_000))
}
