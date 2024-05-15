package main

import (
	"fmt"
	"math/big"
	"os"
	"project-euler/helpers"
)

/*
Let p_n be the nth prime: 2, 3, 5, 7, 11, ..., and let r be the remainder when (p_n - 1)^n + (p_n + 1)^n is divided by p_n^2.
For example, when n = 3, p_3 = 5, and 4^3 + 6^3 = 280 ≡ 5 mod 25.
The least value of n for which the remainder first exceeds 10^9 is 7037.
Find the least value of n for which the remainder first exceeds 10^10.
*/
func main() {
	k := 100_000
	limit := 100 * k
	sieve := helpers.SieveOfEratosthenes(limit)
	primes := []int{}
	for i := 1; i <= limit; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	for n := 1; n <= k; n++ {
		p := primes[n-1]
		rem := nPowerSumMod(int64(p), int64(n))
		if rem.Cmp(big.NewInt(10_000_000_000)) > 0 {
			fmt.Println(n)
			os.Exit(0)
		}
	}
}

func nPowerSumMod(p int64, n int64) *big.Int {
	if n%2 == 0 {
		return big.NewInt(2)
	}

	base := big.NewInt(2)
	bigP := big.NewInt(p)
	base.Mul(base, bigP)
	base.Mul(base, big.NewInt(n))
	base.Mod(base, new(big.Int).Mul(bigP, bigP))
	return base
}
