package main

import (
	"fmt"
	"math/big"
	"os"
	"project-euler/helpers"
	"strconv"
)

/*
The Fibonacci sequence is defined by the recurrence relation:

	F_n = F_(n-1) + F_(n-2), where F_1 = 1 and F_2 = 1.

It turns out that F_541, which contains 113 digits, is the first Fibonacci number for which the last nine digits are 1-9 pandigital (contain all the digits 1 to 9, but not necessarily in order).
And F_2749, which contains 575 digits, is the first Fibonacci number for which the first nine digits are 1-9 pandigital.
Given that F_k  is the first Fibonacci number for which the first nine digits AND the last nine digits are 1-9 pandigital, find k.
*/
func main() {
	f1, f2, fn := big.NewInt(1), big.NewInt(1), big.NewInt(1)
	oneToNine := int64(123456789)
	for i := 3; i <= 10_000_000; i++ {
		fn = fn.Add(f2, f1)

		back := new(big.Int)
		back = back.Mod(fn, big.NewInt(1_000_000_000))
		if helpers.AreDigitPermutations(back.Int64(), oneToNine) {
			fString := fn.String()
			front, _ := strconv.ParseInt(fString[0:min(len(fString), 9)], 10, 64)
			if helpers.AreDigitPermutations(front, oneToNine) {
				fmt.Println(i)
				os.Exit(0)
			}
		}

		f1.Add(f2, big.NewInt(0))
		f2.Add(fn, big.NewInt(0))
	}
}
