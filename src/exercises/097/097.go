package main

import "math/big"

/*
In 2004 there was found a massive non-Mersenne prime which contains 2 357 207 digits: 28433x2^(7830457)+1
*/
func main() {
	modulus := big.NewInt(10_000_000_000)
	exp := big.NewInt(28433)
	for i := 1; i <= 7830457; i++ {
		exp.Mul(exp, big.NewInt(2))
		exp.Mod(exp, modulus)
	}
	exp.Add(exp, big.NewInt(1))
	print(exp.String())
}
