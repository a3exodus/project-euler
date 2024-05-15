package main

import (
	"fmt"
	"project-euler/helpers"
)

/*
The prime factors of 13195 are 5,7,13 and 29. What is the largest prime factor of the number 600851475143?
*/
func compute(n int64) int64 {
	factors := helpers.Factorize(n)
	lastPrime := int64(0)
	for prime, _ := range factors {
		lastPrime = prime
	}
	return lastPrime
}

func main() {
	fmt.Println(compute(600851475143))
}
