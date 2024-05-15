package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
A composite is a number containing at least two prime factors. For example, 15=3x5,9=3x3,12=2x2x3.
There are ten composites below thirty containing precisely two, not necessarily distinct, prime factors:

	4,6,9,10,14,15,21,22,22,25,26

How many composite integers, n<10^8, have precisely two, not necessarily distinct, prime factors?
*/
func compute(n int) int {
	sieve := helpers.SieveOfEratosthenes(int(n / 2))
	primes := []int{}
	for s := range sieve {
		if sieve[s] {
			primes = append(primes, s)
		}
	}

	count := 0
	for _, p := range primes {
		for _, q := range primes {
			if q <= p {
				prod := int64(p) * int64(q)
				if prod <= int64(n) {
					count++
				} else {
					break
				}
			} else {
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(compute(100_000_000))
}
