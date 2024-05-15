package main

import (
	"fmt"
	"project-euler/helpers"
	"strconv"
)

/*
Take the number 192 and multiply it by each of 1, 2, and 3:
192 x 1 = 192
192 x 2 = 384
192 x 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1, 2, 3).
The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1, 2, 3, 4, 5).
What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1, 2, ..., n) where n > 1?
*/
func main() {
	maxPandigital := int64(1)
	for a := 1; a < 10_000; a++ {
		seq := generateSeq(a)
		pandigital, _ := strconv.ParseInt(seq, 10, 64)
		if helpers.AreDigitPermutations(pandigital, 123456789) {
			maxPandigital = max(maxPandigital, pandigital)
		}
	}
	fmt.Println(maxPandigital)
}

func generateSeq(n int) string {
	seq := ""
	multiplier := 1
	for len(seq) < 9 {
		seq += strconv.Itoa(n * multiplier)
		multiplier++
	}
	return seq
}
