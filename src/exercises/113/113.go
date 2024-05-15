package main

import "fmt"

/*
Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.
Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.
We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers below one-thousand (525) are bouncy.
In fact, the least number for which the proportion of bouncy numbers first reaches 50% is 538.

Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the proportion of bouncy numbers is equal to 90%.
Find the least number for which the proportion of bouncy numbers is exactly 99%.
.
*/
func main() {
	numberNonBouncyBelowPower10(100)
}

func numberNonBouncyBelowPower10(power int) int64 {
	N := make([][]int, power)
	d := make([][]int, power)
	i := make([][]int, power)
	for p := 1; p <= power; p++ {
		N[p-1] = make([]int, 10)
		d[p-1] = make([]int, 10)
		i[p-1] = make([]int, 10)
	}
	for n := 0; n <= 9; n++ {
		N[0][n] = 1
		d[0][n] = 1
		i[0][n] = 1
	}

	for p := 2; p <= power; p++ {
		for n := 1; n <= 9; n++ {
			d[p-1][n] = gridRowSum(d, p-2, 0, n)
			i[p-1][n] = gridRowSum(i, p-2, n, 9)
			N[p-1][n] = gridRowSum(d, p-2, 0, n-1) + N[p-2][n] + gridRowSum(i, p-2, n+1, 9)
		}
		d[p-1][0] = 1
		i[p-1][0] = 1
		N[p-1][0] = 1
	}

	//fmt.Println("d")
	//printGrid(d)
	//
	//fmt.Println("i")
	//printGrid(i)
	//
	//fmt.Println("N")
	//printGrid(N)

	sum := int64(0)
	for p := 1; p <= power; p++ {
		for n := 1; n <= 9; n++ {
			sum += int64(N[p-1][n])
		}
	}
	fmt.Println(sum)
	return sum
}

func gridRowSum(grid [][]int, row int, col1 int, col2 int) int {
	sum := 0
	for j := col1; j <= col2; j++ {
		sum += grid[row][j]
	}
	return sum
}

func printGrid(grid [][]int) {
	for i := 1; i <= len(grid); i++ {
		fmt.Println(grid[i-1])
	}
}
