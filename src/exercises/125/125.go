package main

import (
	"fmt"
	"math"
	"project-euler/helpers"
	"strconv"
)

/*
The palindromic number 595 is interesting because it can be written as the sum of consecutive squares:
6² + 7² + 8² + 9² + 10² + 11² + 12².

There are exactly eleven palindromes below one-thousand that can be written as consecutive square sums, and the sum of these palindromes is 4164. Note that 1 = 0² + 1² has not been included as this problem is concerned with the squares of positive integers.

Find the sum of all the numbers less than 10^8 that are both palindromic and can be written as the sum of consecutive squares.
*/
func compute(n int) int64 {
	squares := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		squares = append(squares, i*i)
	}

	consecutivePalindromes := map[int]bool{}
	for i := 0; i < len(squares); i++ {
		for j := i + 2; j < len(squares); j++ {
			sum := subSliceSum(squares, i, j)
			if sum > int64(n) {
				break
			}
			if helpers.IsPalindrome(strconv.FormatInt(sum, 10)) {
				consecutivePalindromes[int(sum)] = true
			}
		}
	}

	sum := int64(0)
	for val, _ := range consecutivePalindromes {
		sum += int64(val)
	}
	return sum
}

func subSliceSum(s []int, i int, j int) int64 {
	sum := int64(0)
	for _, k := range s[i:j] {
		sum += int64(k)
	}
	return sum
}

func main() {
	fmt.Println(compute(100_000_000))
}
