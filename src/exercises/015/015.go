package main

import "fmt"

/*
Starting in the top left corner of a 2x2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
How many such routes are there through a 20x20 grid?
*/
func compute(n int) int64 {
	m := n + 1
	grid := make([][]int64, m)
	for i := range grid {
		grid[i] = make([]int64, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				grid[i][j] = 1
				continue
			}

			// i, i >= 1 here
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][m-1]
}

func main() {
	fmt.Println(compute(20))
}
