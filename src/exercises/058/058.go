package main

import (
	"fmt"
	"math"
	"os"
)

/*
Starting with 1 and spiralling anticlockwise in the following way, a square spiral with side length 7 is formed.

	37 36 35 34 33 32 31
	38 17 16 15 14 13 30
	39 18  5  4  3 12 29
	40 19  6  1  2 11 28
	41 20  7  8  9 10 27
	42 21 22 23 24 25 26
	43 44 45 46 47 48 49

It is interesting to note that the odd squares lie along the bottom right diagonal, but what is more interesting is that 8 out of the 13 numbers lying along both diagonals are prime;
that is, a ratio of 8/13 = 62%.
If one complete new layer is wrapped around the spiral above, a square spiral with side length 9 will be formed.
If this process is continued, what is the side length of the square spiral for which the ratio of primes along both diagonals first falls below 10%?
*/
func main() {
	// let length spiral l be l = 1 + 2r, where r is the 'radius'
	count := 0
	current := 1
	jump := 0
	for r := 1; r <= 100_000; r++ {
		jump += 2
		for i := 1; i <= 4; i++ {
			current += jump
			if isPrimeSimple(int64(current)) {
				count++
			}
		}
		diag := 1 + 4*r
		ratio := float64(count) / float64(diag)
		if ratio < 0.10 {
			fmt.Println(1 + 2*r)
			os.Exit(0)
		}
	}
}

func isPrimeSimple(n int64) bool {
	for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
