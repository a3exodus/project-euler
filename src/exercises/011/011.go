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
In the 20x20 grid in 020.txt, what is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20x20 grid?
*/
func main() {
	file, err := os.Open("011.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n := 20
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
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

	maxProd := 1
	for i := 0; i <= n-4; i++ {
		for j := 0; j <= n-4; j++ {
			prod := maxProductSubGrid(grid, i, j)
			if prod > maxProd {
				maxProd = prod
			}
		}
	}
	fmt.Println(maxProd)
}

func maxProductSubGrid(grid [][]int, row int, col int) int {
	maxProd := 1
	prod1, prod2 := 1, 1
	for i := 0; i < 4; i++ {
		prod1 *= grid[row+i][col+i]
		prod2 *= grid[row+(3-i)][col+i]
	}
	maxProd = max(maxProd, prod1, prod2)

	for i := 0; i < 4; i++ {
		prod1 = 1
		prod2 = 1
		for j := 0; j < 4; j++ {
			prod1 *= grid[row+i][col+j]
			prod2 *= grid[row+j][col+i]
		}
		maxProd = max(maxProd, prod1, prod2)
	}

	return maxProd
}
