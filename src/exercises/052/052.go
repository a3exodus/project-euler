package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
)

/*
It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.
Find the smallest positive integer, x, such that 2x,3x,4x,5x, and 6x, contain the same digits.
*/
func main() {
	n := 6
	for i := 1; i <= 1_000_000; i++ {
		for j := 2; j <= n; j++ {
			prod := i * j
			if !helpers.AreDigitPermutations(int64(i), int64(prod)) {
				break
			}

			if j == n {
				fmt.Println(i)
				os.Exit(0)
			}
		}
	}
}
