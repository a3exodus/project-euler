package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97 and 7.
Similar we can work from right to left: 3797, 379, 37, 3.
Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
*/
func main() {
	n := 1_000_000
	sum := 0
	sieve := helpers.SieveOfEratosthenes(n)

outer:
	for i := 9; i < n; i++ {
		num := strconv.Itoa(i)
		for j := 1; j <= len(num); j++ {
			subNum := num[:j]
			subInt, _ := strconv.Atoi(subNum)
			if !sieve[subInt] {
				continue outer
			}

			subNum2 := num[j-1 : len(num)]
			subInt2, _ := strconv.Atoi(subNum2)
			if !sieve[subInt2] {
				continue outer
			}
		}
		sum += i
	}
	fmt.Println(sum)
}
