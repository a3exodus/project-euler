package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type vector struct {
	x int
	y int
}

/*
Using triangles.txt (right click and 'Save Link/Target As...'), a 27K text file containing the co-ordinates of one thousand "random" triangles, find the number of triangles for which the interior contains the origin.
*/
func main() {
	file, err := os.Open("102.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var vectors []vector
	scanner := bufio.NewScanner(file)
	var line string
	count := 0
	for scanner.Scan() {
		line = scanner.Text()
		numbers := strings.Split(line, `,`)

		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		vectors = append(vectors, vector{x, y})
		x, _ = strconv.Atoi(numbers[2])
		y, _ = strconv.Atoi(numbers[3])
		vectors = append(vectors, vector{x, y})
		x, _ = strconv.Atoi(numbers[4])
		y, _ = strconv.Atoi(numbers[5])
		vectors = append(vectors, vector{x, y})

		if pointInTriangle(vector{0, 0}, vectors[0], vectors[1], vectors[2]) {
			count++
		}
		vectors = nil
	}

	fmt.Println(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func pointInTriangle(pt vector, p1 vector, p2 vector, p3 vector) bool {
	var d1, d2, d3 int
	var hasNeg, hasPos bool

	d1 = sign(pt, p1, p2)
	d2 = sign(pt, p2, p3)
	d3 = sign(pt, p3, p1)

	hasNeg = (d1 < 0) || (d2 < 0) || (d3 < 0)
	hasPos = (d1 > 0) || (d2 > 0) || (d3 > 0)

	return !(hasNeg && hasPos)
}

func sign(p1 vector, p2 vector, p3 vector) int {
	return (p1.x-p3.x)*(p2.y-p3.y) - (p2.x-p3.x)*(p1.y-p3.y)
}
