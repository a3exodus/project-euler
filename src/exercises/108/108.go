package main

import (
	"fmt"
	"os"
)

/*
In the following equation x, y, and n are positive integers.

1/x + 1/y = 1/n

For n = 4 there are exactly three distinct solutions:

1/5 + 1/20 = 1/4
1/6 + 1/12 = 1/4
1/8 + 1/8 = 1/4

What is the least value of n for which the number of distinct solutions exceeds one-thousand?
NOTE: This problem is an easier version of Problem 110; it is strongly advised that you solve this one first.
*/
func main() {
	for n := 1; n <= 1_000_000; n++ {
		if diophantineReciprocals(n) > 1_000 {
			fmt.Println(n)
			os.Exit(0)
		}
	}
}

/*
Unique solutions, so we are looking for (x,y) with x <= y
Rewrite the equation to n(x+y) = xy.
Then
(1) If x > 2n, then xy > 2ny = n(y+y) >= n(x+y), no sols.
(2) If x = 2n, then (2n, 2n) is sol.
(3) If x <= n, then 1/x >= 1n, thus 1/x+1/y > 1/n, no sols.
Remains to look for solutions with x in (n+1, n+2, ..., 2n).
Equation can be rewritten to y = nx/(x-n).
*/

func diophantineReciprocals(n int) int {
	count := 0
	for x := n + 1; x <= 2*n; x++ {
		if n*x%(x-n) == 0 {
			count++
		}
	}
	return count
}
