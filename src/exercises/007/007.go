package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
By listing the first six prime numbers: 2,3,5,7,11 and 13,we can see that the 6th prime is 13.

What is the 10001st prime number?
*/
func compute(n int) int {
	sieve := helpers.SieveOfEratosthenes(1_000_000)
	primes := []int{}
	for s := range sieve {
		if sieve[s] {
			primes = append(primes, s)
		}
	}
	return primes[n-1]
}

func main() {
	fmt.Println(compute(10_001))
}
