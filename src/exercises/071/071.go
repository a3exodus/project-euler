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
By listing the set of reduced proper fractions for d <= 1_000_000 in ascending order of size, find the numerator of the fraction immediately to the left of 3/7.
*/
func compute(f fraction, d int) fraction {
	value := helpers.FractionValue(f)
	var minFraction fraction
	var minDistance = float64(d)
	for i := 1; i <= d; i++ {
		numerator := int(math.Floor(float64(i) * value))
		potentialFraction := fraction{Numerator: numerator, Denominator: i}
		distance := value - helpers.FractionValue(potentialFraction)
		if distance < minDistance && potentialFraction != f {
			if helpers.Gcd(int64(numerator), int64(i)) == 1 {
				minDistance = distance
				minFraction = fraction{numerator, i}
			}
		}
	}

	return minFraction
}

func main() {
	fmt.Println(compute(fraction{3, 7}, 1_000_000).Numerator)
}
