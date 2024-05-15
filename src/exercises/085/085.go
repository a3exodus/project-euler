package main

import (
	"fmt"
	"math"
)

/*
By counting carefully it can be seen that a rectangular grid measuring 3 by 2 contains eighteen rectangles.
Although there exists no rectangular grid that contains exactly two million rectangles, find the area of the grid with the nearest solution.
*/
func main() {
	diff, bestDiff, bestArea := int64(0), int64(2_000_000), 0
	for width := 1; width <= 100; width++ {
		for length := 1; length <= 100; length++ {
			diff = int64(math.Abs(float64(getNumberRectangles(width, length) - 2_000_000)))
			if diff < bestDiff {
				bestDiff = diff
				bestArea = width * length
			}
		}
	}
	fmt.Println(bestArea)
}

func getNumberRectangles(width int, length int) int64 {
	count := int64(0)
	for i := 0; i <= width; i++ {
		for j := 0; j <= length; j++ {
			count += int64(width-i) * int64(length-j)
		}
	}
	return count
}
