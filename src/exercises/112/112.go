package main

import (
	"fmt"
	"strconv"
)

/*
Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.
Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.
We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers below one-thousand (525) are bouncy.
In fact, the least number for which the proportion of bouncy numbers first reaches 50% is 538.

Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the proportion of bouncy numbers is equal to 90%.
Find the least number for which the proportion of bouncy numbers is exactly 99%.
.
*/
func main() {
	fmt.Println(getNForRatio(0.99))
}

func getNForRatio(target float64) int {
	count := 0
	var ratio float64
	for i := 100; i <= 100_000_000; i++ {
		if isBouncy(i) {
			count++
		}
		ratio = float64(count) / float64(i)
		if ratio == target {
			return i
		}
	}
	return -1
}

func isBouncy(n int) bool {
	decreasing, increasing := false, false
	seq := strconv.Itoa(n)
	for i := 0; i < len(seq)-1; i++ {
		digit1 := int(seq[i] - '0')
		digit2 := int(seq[i+1] - '0')

		if digit1 > digit2 {
			decreasing = true
		}

		if digit1 < digit2 {
			increasing = true
		}

		if decreasing && increasing {
			return true
		}
	}
	return false
}
