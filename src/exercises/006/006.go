package main

import "fmt"

/*
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func compute(n int) int64 {
	sum, root := int64(0), int64(0)
	for i := 1; i <= n; i++ {
		sum -= int64(i) * int64(i)
		root += int64(i)
	}
	sum += root * root
	return sum
}

func main() {
	fmt.Println(compute(100))
}
