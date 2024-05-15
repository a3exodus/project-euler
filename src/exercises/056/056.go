package main

import (
	"fmt"
	"math/big"
)

/*
A googol (100^100) is a massive number: one followed by one-hundred zeros; it is almost unimaginably large: one followed by two-hundred zeros.
Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, a^b, where a,b < 100, what is the maximum digital sum?
*/
func main() {
	maxSum := 0
	exp := new(big.Int)
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			exp.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			sum := digitSum(exp)
			maxSum = max(maxSum, sum)
		}
	}
	fmt.Println(maxSum)
}

func digitSum(n *big.Int) int {
	sum := 0
	for _, r := range n.String() {
		sum += int(r - '0')
	}
	return sum
}
