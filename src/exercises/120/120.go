package main

import (
	"fmt"
	"math/big"
)

/*
Let r be the remainder when (a-1)^n + (a+1)^n is divided by a^2.
For example, if a=7 and n=3, then r = 42: 6^3 + 8^3 = 728 = 42 mod 49. And as n varies, so too will r, but for a=7 it turns out that r_max=42.
For 3 <= a <= 1000, find sum_(r_max).
*/
func main() {
	// Let r(a,n) be the remainder of (a-1)^n + (a+1)^n mod a^2.
	// It's quite easy to show that r(a,i) = r(a,i+2a).
	// Just work out (a+1)^(2a+i) = (a+1)^(2a)(a+1)^i congruent (a+1)^i, and similarly one can show it for the second term in r(a,i).

	sum := int64(0)
	for a := 3; a <= 1_000; a++ {
		bigR := rMax(a)
		sum += bigR.Int64()
	}
	fmt.Println(sum)
}

func rMax(a int) *big.Int {
	bigA := big.NewInt(int64(a))
	bigR := big.NewInt(0)
	for n := 1; n <= 2*a; n++ {
		a1 := new(big.Int).Sub(bigA, big.NewInt(1))
		a1 = a1.Exp(a1, big.NewInt(int64(n)), nil)
		a2 := new(big.Int).Add(bigA, big.NewInt(1))
		a2 = a2.Exp(a2, big.NewInt(int64(n)), nil)
		r := new(big.Int).Add(a1, a2)
		r = r.Mod(r, new(big.Int).Mul(bigA, bigA))

		if r.Cmp(bigR) > 0 {
			bigR.Set(r)
		}
	}
	return bigR
}
