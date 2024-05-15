package main

import (
	"fmt"
	"strconv"
)

var sum int64

/*
The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:
- d2d3d4 = 406 is divisible by 2
- d3d4d5 = 063 is divisible by 3
- d4d5d6 = 635 is divisible by 5
- d5d6d7 = 357 is divisible by 7
- d6d7d8 = 572 is divisible by 11
- d7d8d9 = 728 is divisible by 13
- d8d9d10 = 289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.
*/
func main() {
	sum = 0
	generatePerms("", "0123456789")
	fmt.Println(sum)
}

func generatePerms(seq string, numbers string) {
	if numbers == "" {
		pandigital, _ := strconv.ParseInt(seq, 10, 64)
		if (pandigital%1000)%17 == 0 {
			sum += pandigital
		}
		return
	}

	l := len(numbers)
	for i := 0; i < l; i++ {
		r := numbers[i : i+1]

		var num int64
		if len(seq) >= 3 {
			num, _ = strconv.ParseInt(seq[len(seq)-3:], 10, 64)
		}

		if len(seq) == 4 {
			if num%2 != 0 {
				return
			}
		}

		if len(seq) == 5 {
			if num%3 != 0 {
				return
			}
		}

		// d6 = 5 can be proven easily.
		// d4d5d6 is divisible by 5, hence d6 is 0 or 5.
		// If d6 is 0, then d7 - d8 must be divisible by 11, hence d7 = d8 which is not possible.
		if len(seq) == 6 {
			if num%10 != 5 {
				return
			}
		}

		if len(seq) == 7 {
			if num%7 != 0 {
				return
			}
		}

		if len(seq) == 8 {
			if num%11 != 0 {
				return
			}
		}

		if len(seq) == 9 {
			if num%13 != 0 {
				return
			}
		}

		generatePerms(seq+r, numbers[:i]+numbers[i+1:])
	}
}
