package main

import (
	"fmt"
	"math/big"
)

/*
There are exactly ten ways of selecting three from five, 12345:

	123,124,125,134,135,145,234,235,245 and 345.

In combinators we use the notation (5 choose 3) = 10.
In general, (n choose r) = n!/(r! * (n-r)!), where n! = 1*2*...*n, and 0!=1.
It is not until n=23, that a value exceeds one-million, (23 choose 10) = 1144066.
How many, not necessarily distinct, values of (n choose r), 1 <= n <= 100, are greater than 1 million?
*/
func main() {
	million := big.NewInt(1_000_000)
	count := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			val := choose(n, r)
			if val.Cmp(million) > 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func choose(n int, r int) *big.Int {
	val := big.NewInt(1)
	for i := 1; i <= n; i++ {
		val.Mul(val, big.NewInt(int64(i)))
	}
	for i := 1; i <= r; i++ {
		val.Div(val, big.NewInt(int64(i)))
	}
	for i := 1; i <= n-r; i++ {
		val.Div(val, big.NewInt(int64(i)))
	}
	return val
}
