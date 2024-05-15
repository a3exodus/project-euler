package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
The proper divisors of a number are all the divisors excluding the number itself. For example, the proper divisors of 28 are 1, 2, 4, 7, and 14. As the sum of these divisors is equal to 28, we call it a perfect number.

Interestingly the sum of the proper divisors of 220 is 284 and the sum of the proper divisors of 284 is 220, forming a chain of two numbers. For this reason, 220 and 284 are called an amicable pair.

Perhaps less well known are longer chains. For example, starting with 12496, we form a chain of five numbers:
12496 → 14288 → 15472 → 14536 → 14264(→ 12496 → ...)

Since this chain returns to its starting point, it is called an amicable chain.

Find the smallest member of the longest amicable chain with no element exceeding one million.
*/
func main() {
	maxChainLength := 0
	var maxChain map[int]bool
	for i := 1; i < 50_000; i++ {
		chain := amicableChain(i)
		if len(chain) > maxChainLength {
			maxChain = chain
		}
	}

	minValue := 1_000_000
	for value, _ := range maxChain {
		minValue = min(minValue, value)
	}
	fmt.Println(minValue)
}

func amicableChain(n int) map[int]bool {
	chain := map[int]bool{}
	start := n

	for {
		chain[n] = true
		n = helpers.SumDivisors(n)
		if n == 0 || n > 1_000_000 {
			return nil
		}
		_, exists := chain[n]
		if exists {
			if n == start {
				return chain
			}
			return nil
		}
	}
}
