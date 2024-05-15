package main

import (
	"fmt"
	"math/big"
)

/*
2^15 = 32768 and the sum of its digits is 3+2+7+6+8=26. What is the sum of the digits of the number 2^1000?
*/
func compute(n int) int {
	b, p := big.NewInt(2), big.NewInt(int64(n))
	b.Exp(b, p, nil)
	sum := 0
	for _, rune := range b.String() {
		sum += int(rune - '0')
	}
	return sum
}

func main() {
	fmt.Println(compute(1_000))
}
