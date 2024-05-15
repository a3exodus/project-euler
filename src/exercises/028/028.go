package main

import "fmt"

/*
Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows

	21 22 23 24 25
	20  7  8  9 10
	19  6  1  2 11
	18  5  4  3 12
	17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.
What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
*/
func compute(radius int) int64 {
	// let length spiral l be l = 1 + 2r, where r is the 'radius'
	sum := int64(1)
	current := 1
	jump := 0
	for r := 1; r <= radius; r++ {
		jump += 2
		for i := 1; i <= 4; i++ {
			current += jump
			sum += int64(current)
		}
	}
	return sum
}

func main() {
	fmt.Println(compute(500))
}
