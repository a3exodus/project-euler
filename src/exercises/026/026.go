package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
)

/*
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2 = 0.5
1/3 = 0.(3)
1/4 = 0.25
1/5 = 0.2
1/6 = 0.1(6)

Where 0.1(6) means 0.1666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.
Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/
func main() {
	// The length of the repetend (period of the repeating decimal segment) of 1/p is equal to the order of 10 modulo p.
	// Tried biggest such p below one thousand.
	limit := 1000
	sieve := helpers.SieveOfEratosthenes(limit)
	for d := limit; d >= 2; d-- {
		if sieve[d] && ord(d) == d-1 {
			fmt.Println(d)
			os.Exit(0)
		}
	}
}

// https://en.wikipedia.org/wiki/Repeating_decimal#Reciprocals_of_composite_integers_coprime_to_10
func ord(p int) int {
	mod := 1
	for i := 1; i <= p-1; i++ {
		mod *= 10
		mod %= p
		if mod == 1 {
			return i
		}
	}
	return -1
}
