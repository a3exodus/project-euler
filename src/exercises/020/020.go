package main

import (
	"fmt"
	"math/big"
)

/*
Find the sum of the digits in the number 100!
*/
func main() {
	product := big.NewInt(1)
	for i := int64(1); i <= 100; i++ {
		product.Mul(product, big.NewInt(i))
	}
	fmt.Println(product)

	sum := 0
	for _, rune := range product.String() {
		sum += int(rune - '0')
	}
	fmt.Println(sum)
}
