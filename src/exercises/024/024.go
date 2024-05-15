package main

import (
	"fmt"
	"os"
)

var count int = 0

/*
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

	012 021 102 120 201 210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/
func main() {
	generatePerms("", "0123456789")
}

func generatePerms(seq string, numbers string) {
	if numbers == "" {
		count++
		if count == 1_000_000 {
			fmt.Println(seq)
			os.Exit(0)
		}
		return
	}

	l := len(numbers)
	for i := 0; i < l; i++ {
		r := numbers[i : i+1]
		generatePerms(seq+r, numbers[:i]+numbers[i+1:])
	}
}
