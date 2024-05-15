package main

import (
	"fmt"
	"math/big"
)

/*
Find the last ten digits of the series, 1^1 + 2^2 + ... + 1000^1000.
*/
func compute(n int) *big.Int {
	sum := new(big.Int)
	modulo := big.NewInt(10_000_000_000)
	for i := 1; i <= n; i++ {
		bigI := big.NewInt(int64(i))
		sum.Add(sum, new(big.Int).Exp(bigI, bigI, nil))
	}
	out := sum.Mod(sum, modulo)
	return out
}

func main() {
	fmt.Println(compute(1_000))
}
