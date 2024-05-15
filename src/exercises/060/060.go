package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
	"strconv"
)

var sieve []bool

/*
The primes 3,7,109,673 are quite remarkable.  By taking any two primes and concatenating them in any order the result will always be prime.
For example, taking 7 and 109, both 7109 and 107 are prime. The sum of these four primes, 792, represents the lowest sum for a set of four primes with this property.
Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
*/
func main() {

	n := 10_000
	limit := n * n
	sieve = helpers.SieveOfEratosthenes(limit)
	primes := []int{}
	for s := range sieve {
		if sieve[s] {
			primes = append(primes, s)
		}
	}

	for _, p := range primes {
		if p < n {
			for _, q := range primes {
				if q < n && p < q {
					if isConcatPrime(p, q) {
						for _, r := range primes {
							if r < n && q < r {
								if isConcatPrime(p, r) && isConcatPrime(q, r) {
									for _, s := range primes {
										if s < n && r < s {
											if isConcatPrime(p, s) && isConcatPrime(q, s) && isConcatPrime(r, s) {
												for _, t := range primes {
													if t < n && r < t {
														if isConcatPrime(p, t) && isConcatPrime(q, t) && isConcatPrime(r, t) && isConcatPrime(s, t) {
															fmt.Println(p + q + r + s + t)
															os.Exit(0)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

}

func isConcatPrime(n int, m int) bool {
	concat, _ := strconv.ParseInt(strconv.Itoa(n)+strconv.Itoa(m), 10, 64)
	if sieve[concat] {
		concat2, _ := strconv.ParseInt(strconv.Itoa(m)+strconv.Itoa(n), 10, 64)
		return sieve[concat2]
	}
	return false
}
