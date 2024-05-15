package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
Euler's totient function, phi(n)  [sometimes called the phi function], is used to determine the number of positive numbers less than or equal to n which are relatively prime to n.
For example, as 1,2,4,5,7 and 8, are all less than nine and relatively prime to nine, phi(9) = 6.

Interestingly, phi(87109) = 79180, and it can be seen that 87109 is a permutation of 79180.

Find the value of n, 1 < n < 10^7, for which phi(n) is a permutation of n and the ratio n/phi(n) produces a minimum.
*/
func main() {
	minRatio, minN := 3.0, 0
	for n := 2; n <= 10_000_000; n++ {
		phiN := helpers.Phi(n)
		ratio := float64(n) / float64(phiN)
		if ratio < minRatio {
			if helpers.AreDigitPermutations(int64(n), int64(phiN)) {
				minRatio = ratio
				minN = n
			}
		}
	}
	fmt.Println(minN)
}
