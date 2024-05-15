package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
)

/*
It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9=7+2*1^2
15=7+2*2^2
21=3+2*3^2
25=7+2*3^2
27=19+2*2^2
33=31+2*1^2

It turns out that the conjecture was false.
What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
*/
func main() {
	n := 10_000
	sieve := helpers.SieveOfEratosthenes(n)

	goldbach := map[int]bool{}
	for s := range sieve {
		if sieve[s] {
			for i := 1; i < n; i++ {
				sum := s + 2*i*i
				goldbach[sum] = true
			}
		}
	}

	// check if odd composite number not in goldbach
	for i := 3; i < n; i += 2 {
		if !sieve[i] {
			_, found := goldbach[i]
			if !found {
				fmt.Println(i)
				os.Exit(0)
			}
		}
	}
}
