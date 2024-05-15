package main

import (
	"fmt"
	"os"
	"project-euler/helpers"
	"strconv"
)

/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.
What is the largest n-digit pandigital prime that exists?
*/
func main() {

	m := 100_000_000
	sieve := helpers.SieveOfEratosthenes(m)
	pandigital := "123456789"

	for i := m; i > 2; i-- {
		if sieve[i] {
			num := strconv.Itoa(i)
			subPandigital, _ := strconv.ParseInt(pandigital[0:len(num)], 10, 64)
			if helpers.AreDigitPermutations(int64(i), subPandigital) {
				fmt.Println(i)
				os.Exit(0)
			}
		}
	}
}
