package main

import (
	"fmt"
	"math"
	"project-euler/helpers"
)

/*
The smallest number expressible as the sum of a prime square, prime cube, and prime fourth power is 28. In fact, there are exactly four numbers below fifty that can be expressed in such a way:

	28 = 2^2 + 2^3 + 2^4
	33 = 3^2 + 2^3 + 2^4
	49 = 5^2 + 2^3 + 2^4
	47 = 2^2 + 3^2 + 2^4

How many numbers below fifty million can be expressed as the sum of a prime square, prime cube, and prime fourth power?
*/
func compute(n int) int {
	m := int(math.Sqrt(float64(n)))
	sieve := helpers.SieveOfEratosthenes(m)
	primes := []int{}
	for i := 1; i <= m; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	targets := map[int]bool{}
	for _, p := range primes {
		for _, q := range primes {
			for _, r := range primes {
				potential := p*p + q*q*q + r*r*r*r
				if potential < n {
					targets[potential] = true
				} else {
					break
				}
			}
		}
	}
	return len(targets)
}

func main() {
	fmt.Println(compute(50_000_000))
}
