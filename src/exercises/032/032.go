package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.
The product 7254 is unusual, as the identity, 39x186=7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.
Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
*/
func main() {

	pandigital := int64(123456789)
	validProds := map[int]bool{}
	for a := 1; a < 10_000; a++ {
		for b := a + 1; b < 10_000; b++ {
			prod := a * b

			// skip unnecessary iterations
			if prod >= 10_000 {
				break
			}

			concat, _ := strconv.ParseInt(strconv.Itoa(a)+strconv.Itoa(b)+strconv.Itoa(prod), 10, 64)
			if helpers.AreDigitPermutations(concat, pandigital) {
				validProds[prod] = true
			}
		}
	}

	sum := 0
	for validProd := range validProds {
		sum += validProd
	}
	fmt.Println(sum)
}
