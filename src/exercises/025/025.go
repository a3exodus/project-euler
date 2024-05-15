package main

import (
	"fmt"
	"math/big"
)

/*
What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/
func compute(numDigits int) int {
	f1, f2, fn := big.NewInt(1), big.NewInt(1), big.NewInt(1)
	n := 10_000
	for i := 3; i <= n; i++ {
		fn = fn.Add(f2, f1)
		if len(fn.String()) >= numDigits {
			return i
			break
		}
		f1.Add(f2, big.NewInt(0))
		f2.Add(fn, big.NewInt(0))
	}
	return -1
}

func main() {
	fmt.Println(compute(1_000))
}
