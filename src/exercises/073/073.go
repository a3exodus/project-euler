package main

import (
	"fmt"
	"math"
	"project-euler/helpers"
)

type fraction = helpers.Fraction

/*
Consider the fraction, n/d, where n and d are positive integers.
If n < d and gcd(n,d) = 1, it is called a poper reduced fraction.
How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for d <= 12_000?
?
*/
func compute(left fraction, right fraction, d int) int {
	count := 0
	leftValue := helpers.FractionValue(left)
	rightValue := helpers.FractionValue(right)
	for i := 2; i <= d; i++ {
		lowerBound := int(math.Ceil(float64(i) * leftValue))
		upperBound := int(math.Ceil(float64(i) * rightValue))
		for j := lowerBound; j <= upperBound; j++ {
			potentialFraction := fraction{Numerator: j, Denominator: i}
			potentialFractionValue := helpers.FractionValue(potentialFraction)
			if leftValue < potentialFractionValue && potentialFractionValue < rightValue {
				if helpers.Gcd(int64(j), int64(i)) == 1 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	count := compute(fraction{Numerator: 1, Denominator: 3}, fraction{Numerator: 1, Denominator: 2}, 12_000)
	fmt.Println(count)
}
