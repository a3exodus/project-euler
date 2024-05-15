package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
)

/*
The first two consecutive numbers to have two distinct prime factors are:
14 = 2x7
15 = 3x5
The first three consecutive numbers to have three distinct prime factors are:
644 = 2^2x7x23
645 = 3x5x43
646 = 2x17x19
Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?
*/
func main() {
	consecutive := 4
	for i := 1; i <= 1_000_000; i++ {
		for j := 0; j < consecutive; j++ {
			num := i + j
			if len(helpers.Factorize(int64(num))) != consecutive {
				break
			}

			if num == i+consecutive-1 {
				fmt.Println(i)
				os.Exit(0)
			}
		}
	}
}
