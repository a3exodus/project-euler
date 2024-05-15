package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
Consider the fraction, n/d, where n and d are positive integers.
If n < d and gcd(n,d) = 1, it is called a poper reduced fraction.
How many elements would be contained in the set of reduced proper fractions for d <= 1_000_000?
*/
func compute(d int) int64 {
	total := int64(0)
	for i := 2; i <= d; i++ {
		total += int64(helpers.Phi(i))
	}
	return total
}

func main() {
	fmt.Println(compute(1_000_000))
}
