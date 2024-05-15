package main

import (
	"fmt"
	"strconv"
)

var count int = 0

/*
It is possible to write five as a sum in exactly six different ways:

	4 + 1
	3 + 2
	3 + 1 + 1
	2 + 2 + 1
	2 + 1 + 1 + 1
	1 + 1 + 1 + 1 + 1

How many different ways can one hundred be written as a sum of at least two positive integers?
*/
func main() {
	n := 100
	generatePerms(n, "", n-1)
	fmt.Println(count)
}

func generatePerms(remainder int, seq string, maxterm int) {
	if remainder < 0 {
		return
	}

	if remainder == 0 {
		count++
	}

	for i := 1; i <= maxterm; i++ {
		generatePerms(remainder-i, seq+strconv.Itoa(i), i)
	}
}
