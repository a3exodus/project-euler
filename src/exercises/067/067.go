package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is

3
7 4
2 4 6
8 5 9 3

That is, 3+7+4+9=23.
Find the maximum total from top to bottom of the triangle in 067.txt.
*/
func main() {
	file, err := os.Open("067.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n := 100
	grid := make([][]int, n)
	path := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
		path[i] = make([]int, n)
	}

	scanner := bufio.NewScanner(file)
	row, col := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ` `)
		for _, num := range numbers {
			grid[row][col], _ = strconv.Atoi(num)
			col++
		}
		row++
		col = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	path[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			var maxPath int
			if j == 0 {
				maxPath = path[i-1][0]
			} else {
				maxPath = max(path[i-1][j], path[i-1][j-1])
			}
			path[i][j] = maxPath + grid[i][j]
		}
	}

	maxPath := 0
	for j := 0; j < n; j++ {
		maxPath = max(maxPath, path[n-1][j])
	}
	fmt.Println(maxPath)
}
