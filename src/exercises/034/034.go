package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
Find the sum of all numbers which are equal to the sum of the factorial of their digits.
Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/
func main() {
	sum := 0
	for i := 3; i <= 10_000_000; i++ {
		dfsum := digitFactorialSum(i)
		if i == dfsum {
			sum += i
		}
	}
	fmt.Println(sum)
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
