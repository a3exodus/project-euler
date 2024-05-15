package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
func main() {
	sum := 0
	for i := 1; i <= 1_000_000; i++ {
		if helpers.IsPalindrome(strconv.Itoa(i)) {
			binary := strconv.FormatInt(int64(i), 2)
			if helpers.IsPalindrome(binary) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}
