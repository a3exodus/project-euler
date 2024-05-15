package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
Find the number of integers 1 < n < 10^7, for which n and n+1 have the same number of positive divisors. For example, 14 has the positive divisors 1,2,7,14 while 15 has 1,3,5,15.
*/
func compute(n int) int {
	numDivs := make([]int64, n)
	count := 0
	numDivs[0] = 1
	for i := 2; i < n; i++ {
		numDivs[i-1] = helpers.NumDivisors(int64(i))
		if numDivs[i-1] == numDivs[i-2] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(compute(10_000_000))
}
