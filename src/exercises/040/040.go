package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
An irrational decimal fraction is created by concatenating the positive integers:

0.12345678910111213...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1xd10xd100x...d_{1_000_000}
*/
func main() {
	prod := 1
	seqLength := 0
	target := 1
	for i := 1; i <= 1_000_000; i++ {
		num := strconv.Itoa(i)
		length := len(num)
		if seqLength+length >= target {
			diff := target - seqLength
			digit := int(num[diff-1] - '0')
			target *= 10
			prod *= digit
			if target > 1_000_000 {
				fmt.Println(prod)
				os.Exit(0)
			}
		}
		seqLength += len(num)
	}

}
