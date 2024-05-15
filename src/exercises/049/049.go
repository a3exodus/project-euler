package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.
There are no arithmetic sequences made up of three 1-,2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.
What 12-digit number do you form by concatenating the three terms in this sequence?
*/
func main() {

	sieve := helpers.SieveOfEratosthenes(10_000)

	for i := 1000; i <= 3000; i++ {
		if sieve[i] {
			j := i + 3330
			if sieve[j] {
				k := j + 3330
				if sieve[k] {
					if helpers.AreDigitPermutations(int64(i), int64(j)) {
						concat, _ := strconv.ParseInt(strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k), 10, 64)

						if i == 1487 {
							continue
						}
						fmt.Println(concat)
					}
				}
			}
		}
	}

}
