package main

import (
	"fmt"
	"math/big"
)

/*
It is possible to show that the square root of two can be expressed as an infinite continued fraction.

	root(2) = 1 + 1/(2+1/(2+...))

By expanding this for the first four iterations, we get:

	1 + 1/2 = 3/2 = 1.5
	1 + 1/(2+1/2) = 7/5 = 1.4
	1 + 1/(2+1/(1+1/2)) = 17/12 = 1.41666
	1 + 1/(2+1/(1+1/(1+1/2))) = 41/29 = 1.41379

The next three expansions are 99/70, 239/169 and 577/408, but the eighth expansion, 1393/985,
is the first example where the number of digits in the numerator exceeds the number of digits in the denominator.

In the first one-thousand expansions, how many fractions contain a numerator with more digits than the denominator?
*/
func main() {
	numerator := big.NewInt(1)
	denominator := big.NewInt(1)
	tmp := new(big.Int)

	count := 0
	for i := 1; i <= 1_000; i++ {
		tmp.Set(denominator)
		denominator.Add(denominator, numerator)
		numerator.Add(denominator, tmp)
		if len(numerator.String()) > len(denominator.String()) {
			count++
		}
	}
	fmt.Println(count)
}
