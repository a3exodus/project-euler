package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Using network.txt (right click and 'Save Link/Target As...'), a 6K text file containing a network with forty vertices, and given in matrix form, find the maximum saving which can be achieved by removing redundant edges whilst ensuring that the network remains connected.
*/
func main() {
	file, err := os.Open("107.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n := 40
	edges := make([][]int, n)
	for i := range edges {
		edges[i] = make([]int, n)
	}

	scanner := bufio.NewScanner(file)
	row, col := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, `,`)
		for _, num := range numbers {
			edges[row][col], _ = strconv.Atoi(num)
			col++
		}
		row++
		col = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	networkValue := totalEdgesValue(edges)
	minNetworkValue := primsAlgorithm(edges)
	fmt.Println(networkValue, minNetworkValue, networkValue-minNetworkValue)
}

// https://en.wikipedia.org/wiki/Prim%27s_algorithm
func primsAlgorithm(edges [][]int) int {
	n := len(edges)

	// Step 1
	C := map[int]int{}
	E := map[int]int{}
	for i := 1; i <= n; i++ {
		C[i] = math.MaxInt32
		E[i] = 0
	}

	// Step 2
	// Initialize an empty forest F and a set Q of vertices that have not yet been included in F (initially, all vertices).
	F := []int{}
	Q := map[int]bool{}
	for i := 1; i <= n; i++ {
		Q[i] = true
	}

	sum := 0
	// Step 3: Repeat the following steps until Q is empty
	for len(Q) > 0 {
		// a: Find and remove a vertex v from Q having the minimum possible value of C[v]
		v := getMinVertexQ(Q, C)
		delete(Q, v)

		// b: Add v to F
		F = append(F, v)
		if len(F) > 1 {
			sum += C[v]
		}

		// c: Loop over the edges vw connecting v to other vertices w.
		// For each such edge, if w still belongs to Q and vw has smaller weight than C[w], perform the following steps:
		for w := 1; w <= n; w++ {
			weight := edges[v-1][w-1]
			if weight > 0 {
				if Q[w] && weight < C[w] {
					// i: Set C[w] to the cost of edge vw
					C[w] = weight
					// ii: Set E[w] to point to edge vw.
					E[w] = v
				}
			}
		}
	}

	return sum
}

func getMinVertexQ(Q map[int]bool, C map[int]int) int {
	minVertex := 1
	minDistance := math.MaxInt32
	for v, exists := range Q {
		if exists {
			if C[v] < minDistance {
				minVertex = v
				minDistance = C[v]
			}
		}
	}
	return minVertex
}

func totalEdgesValue(edges [][]int) int {
	n := len(edges)
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			sum += edges[i][j]
		}
	}
	return sum
}
