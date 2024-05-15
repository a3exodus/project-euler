package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
	"strconv"
)

var count int = 0
var sieve []bool

/*
It is possible to write ten as the sum of primes in exactly five different ways:

	7 + 3
	5 + 5
	5 + 3 + 2
	3 + 2 + 2 + 2
	2 + 2 + 2 + 2 + 2

What is the first value which can be written as the sum of primes in over five thousand different ways?
*/
func main() {
	n := 100
	sieve = helpers.SieveOfEratosthenes(n)
	for i := 1; i < n; i++ {
		count = 0
		generatePerms(i, "", i-1)
		if count > 5_000 {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}

func generatePerms(remainder int, seq string, maxterm int) {
	if remainder < 0 {
		return
	}

	if remainder == 0 {
		count++
	}

	for i := 2; i <= maxterm; i++ {
		if sieve[i] {
			generatePerms(remainder-i, seq+strconv.Itoa(i), i)
		}
	}
}
