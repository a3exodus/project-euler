package main

import "fmt"

/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
*/
func main() {

	count := 0
	for i := 1; i <= 999; i++ {
		count += len(digitToString(i))
	}
	count += len("onethousand")
	fmt.Println(count)
}

func digitToString(n int) string {
	if n == 0 {
		return ""
	}
	twentyOut := [...]string{"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	if n < 20 {
		return twentyOut[n-1]
	}

	if n >= 100 {
		hundred := n / 100
		ten := n % 100
		if ten == 0 {
			return twentyOut[hundred-1] + "hundred"
		}
		return twentyOut[hundred-1] + "hundredand" + digitToString(n%100)
	} else {
		ten := n / 10
		tenOut := [...]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
		return tenOut[ten-2] + digitToString(n%10)
	}
}
