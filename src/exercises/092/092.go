package main

import (
	"fmt"
	"strconv"
)

/*
A number chain is created by continuously adding the square of the digits in a number to form a new number until it has been seen before.

For example,

	44 => 32 => 13 => 10 => 1 => 1
	85 => 89 => 145 => 42 => 20 => 4 => 16 => 37 => 58 => 89

Therefore any chain that arrives at 1 or 89 will become stuck in an endless loop. What is most amazing is that EVERY starting number will eventually arrive at 1 or 89.
How many starting numbers below ten million will arrive at 89
*/
func main() {
	n := 100
	initialSeq := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		initialSeq[i] = arrivesAt89(i)
	}

	count := 0
	for i := 1; i < 10_000_000; i++ {
		if arrivesAt89Fast(i, initialSeq) {
			count++
		}
	}
	fmt.Println(count)
}

func arrivesAt89(n int) bool {
	if n == 1 {
		return false
	}
	if n == 89 {
		return true
	}

	sumDigitsSq := sumDigitsSquared(n)
	return arrivesAt89(sumDigitsSq)
}

func arrivesAt89Fast(n int, initialSeq []bool) bool {
	if n < len(initialSeq) {
		return initialSeq[n]
	}

	sumDigitsSq := sumDigitsSquared(n)
	return arrivesAt89(sumDigitsSq)
}

func sumDigitsSquared(n int) int {
	num := strconv.Itoa(n)
	sum := int(0)
	for _, r := range num {
		digit := int(r - '0')
		sum += digit * digit
	}
	return sum
}
