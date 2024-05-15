package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
The number 145 is well known for the property that the sum of the factorial of its digits is equal to 145:
1! + 4! + 5! = 1 + 24 + 120 = 145.

Perhaps less well known is 169, in that it produces the longest chain of numbers that link back to 169; it turns out that there are only three such loops that exist:
169 → 363601 → 1454 → 169
871 → 45361 → 871
872 → 45362 → 872

It is not difficult to prove that EVERY starting number will eventually get stuck in a loop. For example,
69 → 363600 → 1454 → 169 → 363601 (→ 1454)
78 → 45360 → 871 → 45361(→ 871)
540 → 145 (→ 145)

Starting with 69 produces a chain of five non-repeating terms, but the longest non-repeating chain with a starting number below one million is sixty terms.

How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?
*/
func main() {
	count := 0
	for i := 1; i < 1_000_000; i++ {
		chain := factorialChain(i)
		if chain == 60 {
			count++
		}
	}
	fmt.Println(count)
}

func factorialChain(n int) int {
	chain := map[int]bool{}

	for {
		chain[n] = true
		n = digitFactorialSum(n)
		_, exists := chain[n]
		if exists {
			return len(chain)
		}
	}
}

func digitFactorialSum(n int) int {
	num := strconv.Itoa(n)
	sum := 0
	for _, r := range num {
		digit := int(r - '0')
		sum += helpers.Factorial(digit)
	}
	return sum
}
