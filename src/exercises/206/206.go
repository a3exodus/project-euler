package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0, where each “_” is a single digit.
*/
func main() {
	// n = m^2 = 1_2_3_4_5_6_7_8_9_0
	// 2, 5 divides n => 2, 5 divides m => m = 10k
	// find k such that (10k)^2 = 100k^2 = 1_2_3_4_5_6_7_8_900 iff k^2 = 1_2_3_4_5_6_7_8_9

	// After calculating the squares of 1,2,...,999, we see that k = 1000i + j, where j in threes
	threes := []int{}
	for i := 1; i < 999; i++ {
		square := i * i % 1_000
		if 800 <= square && square <= 899 && square%10 == 9 {
			threes = append(threes, i)
		}
	}

	for i := int64(100_000); i < 1_000_000; i++ {
		for _, t := range threes {
			k := 1000*i + int64(t)
			square := k * k
			if IsConcealedSquare(square) {
				m := 10 * k
				fmt.Println(m)
				os.Exit(0)
			}
		}
	}
}

func IsConcealedSquare(square int64) bool {
	num := strconv.FormatInt(square, 10)
	l := len(num)
	if l != 17 {
		return false
	}

	target := "1_2_3_4_5_6_7_8_9"
	for i := 0; i < 17; i++ {
		if i%2 == 0 {
			targetDigit := int(target[i] - '0')
			digit := int(num[i] - '0')
			if digit != targetDigit {
				return false
			}
		}
	}
	return true
}
