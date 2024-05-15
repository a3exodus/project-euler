package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009=91x99.
Find the largest palindrome made from the product of two 3 digit numbers.
*/
func main() {
	max := 1
	for a := 100; a <= 999; a++ {
		for b := a; b <= 999; b++ {
			prod := a * b
			if prod > max {
				if helpers.IsPalindrome(strconv.Itoa(prod)) {
					max = prod
				}
			}
		}
	}
	fmt.Println(max)
}
