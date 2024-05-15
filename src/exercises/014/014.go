package main

import "fmt"

/*
The following iterative sequence is defined for the set of positive integers:

	n => n / 2 if n even
	n => 3n + 1 if n odd

Using the rule above and starting with 13, we generate the following sequence:

	13 => 40 => 20 => 10 => 5 => 16 => 8 => 4 => 2 => 1.

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?
*/
func main() {
	chainLength, maxLength, maxNum := 0, 0, 0
	for i := 1; i < 1_000_000; i++ {
		chainLength = getCollatzChainLength(int64(i))
		if chainLength > maxLength {
			maxLength = chainLength
			maxNum = i
		}
	}
	fmt.Println(maxNum)
}

func getCollatzChainLength(n int64) int {
	if n == 1 {
		return 1
	}

	// n > 1 here
	if n%2 == 0 {
		return 1 + getCollatzChainLength(n/2)
	} else {
		return 1 + getCollatzChainLength(3*n+1)
	}
}
