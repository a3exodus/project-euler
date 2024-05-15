package main

import (
	"fmt"
	"math/big"
)

/*
The 5-digit number, 16807 = 7^5, is also a fifth power. Similarly, the 9-digit number, 134217728 = 8^9, is a ninth power.
How many n-digit positive integers exist which are also an nth power?
*/
func main() {
	count := 0
	val := new(big.Int)
	for b := 1; b < 99; b++ {
		for n := 1; n < 99; n++ {
			a := val.Exp(big.NewInt(int64(b)), big.NewInt(int64(n)), nil)
			if len(a.String()) == int(n) {
				count++
			}
		}
	}
	fmt.Println(count)
}
