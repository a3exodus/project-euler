package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
)

/*
The cube, 41063625 (345^3), can be permuted to produce two other cubes: 56623104 (384^3) and 66430125 (405^3). In fact,
41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are cube.
*/
func main() {

	n := 10_000

	cubes := []int64{}
	for i := 1; i <= n; i++ {
		cube := int64(i) * int64(i) * int64(i)
		cubes = append(cubes, cube)
	}

	for _, cube := range cubes {
		count := 0
		for _, permutation := range cubes {
			if cube <= permutation {
				if helpers.AreDigitPermutations(cube, permutation) {
					count++
					if count == 5 {
						fmt.Println(cube)
						os.Exit(0)
					}
				}
			}
		}
	}
}
