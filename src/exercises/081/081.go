package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"project-euler/helpers"
	"strconv"
	"strings"
)

type Edge = helpers.Edge

/*
Find the minimal path sum from the top left to the bottom right by only moving right and down in matrix.txt (right click and "Save Link/Target As..."), a 31K text file containing an 80 by 80 matrix.
*/
func main() {
	file, err := os.Open("081.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n := 80
	nodes := make([][]int, n)
	for i := range nodes {
		nodes[i] = make([]int, n)
	}

	scanner := bufio.NewScanner(file)
	row, col := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, `,`)
		for _, num := range numbers {
			nodes[row][col], _ = strconv.Atoi(num)
			col++
		}
		row++
		col = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	G := map[string][]Edge{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			v := "(" + strconv.Itoa(i) + "," + strconv.Itoa(j) + ")"
			G[v] = make([]Edge, 0)
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			v := "(" + strconv.Itoa(i) + "," + strconv.Itoa(j) + ")"

			x := 1
			indexX := i + x - 1
			if indicesInBounds(indexX, n) {
				w := "(" + strconv.Itoa(i+x) + "," + strconv.Itoa(j) + ")"
				e := Edge{Vertex: w, Weight: nodes[indexX][j-1]}
				G[v] = append(G[v], e)
			}

			y := 1
			indexY := j + y - 1
			if indicesInBounds(indexY, n) {
				w := "(" + strconv.Itoa(i) + "," + strconv.Itoa(j+y) + ")"
				e := Edge{Vertex: w, Weight: nodes[i-1][indexY]}
				G[v] = append(G[v], e)
			}
		}
	}

	source, target := "(1,1)", "(80,80)"
	dist := helpers.DijkstrasAlgorithm(G, source)
	fmt.Println(nodes[0][0] + dist[target])
}

func indicesInBounds(i int, n int) bool {
	return 0 <= i && i < n
}
