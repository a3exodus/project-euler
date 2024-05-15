package main

import (
	"fmt"
	"math/big"
	"project-euler/helpers"
)

/*
If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.
Not all numbers produce palindromes so quickly. For example,

	349 + 943 = 1292
	1292 + 2921 = 4212
	4213 + 3124 = 7337

that is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome.
A number that never forms a palindrome through the reverse and add process is called a Lychrel number.
How many Lychrel numbers are there below ten-thousand?
*/
func main() {
	//isLychrel(big.NewInt(10677))
	count := 0
	for i := 1; i < 10_000; i++ {
		j := big.NewInt(int64(i))
		if isLychrel(j) {
			count++
		}
	}
	fmt.Println(count)
}

func isLychrel(n *big.Int) bool {
	for i := 1; i < 50; i++ {
		n = n.Add(n, reverse(n))
		if helpers.IsPalindrome(n.String()) {
			return false
		}
	}
	return true
}

func reverse(n *big.Int) *big.Int {
	var reverse string
	for _, v := range n.String() {
		reverse = string(v) + reverse
	}

	out := new(big.Int)
	out, _ = out.SetString(reverse, 10)
	return out
}
