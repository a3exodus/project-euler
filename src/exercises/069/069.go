package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
Euler's totient function, phi(n) [sometimes called the phi function], is defined as the number of positive integers not exceeding n which are relatively prime to n.
For example, as 1, 2, 4, 5, 7, and 8, are all less than or equal to nine and relatively prime to nine, phi(9)=6.

n Relatively Prime! ¢(n)| n/¢(n)
2 1 1 2
3 1,2 2 1.5
4 1,3 2 2
5 1,2,3,4 4 1.25
6 1,5 2 3
7 1,2,3,4,5,6 6 1.1666.
8 1,3,5,7 4 2
9 1,2,4,5,7,8 6 1.5
10 1,3,7,9 4 25

It can be seen that n = 6 produces a maximum n/phi(n) for n < 10.
Find the value of n < 1000000 for which n/phi(n) is a maximum.
*/
func main() {
	maxRatio, maxN := 0.0, 0
	for n := 2; n <= 1_000_000; n++ {
		phiN := helpers.Phi(n)
		ratio := float64(n) / float64(phiN)
		if ratio > maxRatio {
			maxRatio = ratio
			maxN = n
		}
	}
	fmt.Println(maxN)
}
