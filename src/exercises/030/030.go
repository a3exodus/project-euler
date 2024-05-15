package main

import (
	"fmt"
	"strconv"
)

/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
1634=1^4+6^4+3^4+4^4
8208=8^4+2^4+0^4+4^4
9474=9^4+4^4+7^4+4^4
Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/
func main() {
	sum := 0
	for i := 2; i <= 1_000_000; i++ {
		if i == digitPowerSum(i) {
			sum += i
		}
	}
	fmt.Print(sum)
}

func digitPowerSum(n int) int {
	sum := 0
	for _, r := range strconv.Itoa(n) {
		digit := int(r - '0')
		sum += digit * digit * digit * digit * digit
	}
	return sum
}
