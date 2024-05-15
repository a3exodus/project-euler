package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
The prime 41, can be written as the sum of six consecutive primes:

	41 = 2 + 3 + 5 + 7 + 11 + 13.

This is the longest sum of consecutive primes that adds to a prime below one-hundred.
The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.
Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/
func main() {
	n := 1_000_000
	sieve := helpers.SieveOfEratosthenes(n)
	primes := []int{}
	for s := range sieve {
		if sieve[s] {
			primes = append(primes, s)
		}
	}

	longestChain, longestTarget := 0, 1
	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			if j-i > longestChain {
				sum := subSliceSum(primes, i, j)
				if sum > n {
					break
				}
				if sieve[sum] {
					longestChain = j - i
					longestTarget = sum
				}
			}
		}
	}
	fmt.Println(longestTarget)
}

func subSliceSum(s []int, i int, j int) int {
	sum := 0
	for _, k := range s[i:j] {
		sum += k
	}
	return sum
}
