package main

import (
	"fmt"
	"math"
	"math/big"
)

/*
It is well known that if the square root of a natural number is not an integer, then it is irrational.
The decimal expansion of such square roots is infinite without any repeating pattern at all.
The square root of two is 1.4142135623..., and the digital sum of the first one hundred decimal digits is 475.
For the first one hundred natural numbers, find the total of the digital sums of the first one hundred decimal digits for all the irrational square roots.
*/
func main() {
	sum := 0
	for i := 2; i <= 100; i++ {
		root := int(math.Sqrt(float64(i)))
		if i != root*root {
			sum += sumSqrt(i)
		}
	}
	fmt.Println(sum)
}

func sumSqrt(sqrt int) int {
	precision := uint(10_000)
	numFloat := big.NewFloat(float64(sqrt)).SetPrec(1000)
	root := numFloat.Sqrt(numFloat).SetPrec(precision)
	for i := 1; i <= 99; i++ {
		root = root.Mul(root, big.NewFloat(10))
	}
	num, _ := root.Int(nil)
	fmt.Println(num)

	return sumDigits(num.String())
}

func sumDigits(seq string) int {
	sum := 0
	for _, r := range seq {
		sum += int(r - '0')
	}
	return sum
}
