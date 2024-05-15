package main

import (
	"fmt"
	"project-euler/helpers"
)

type fraction = helpers.Fraction

/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 4/8, which is correct, is obtained by cancelling the 9s.
We shall consider fractions like, 30/50=3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.
If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/
func main() {

	fractions := []fraction{}
	for denom := 10; denom <= 99; denom++ {
		for num := 10; num < denom; num++ {
			f := fraction{num, denom}
			value := helpers.FractionValue(f)

			if num%10 == denom/10 {
				if helpers.FractionValue(fraction{num / 10, denom % 10}) == value {
					fractions = append(fractions, f)
				}
			}
		}
	}

	num := 1
	denom := 1
	for _, f := range fractions {
		num *= f.Numerator
		denom *= f.Denominator
	}

	fmt.Println(denom / int(helpers.Gcd(int64(num), int64(denom))))
}
