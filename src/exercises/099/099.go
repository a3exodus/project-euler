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
Using 099.txt, a 22K text file containing one thousand lines with a base/exponent pair on each line, determine which line number has the greatest numerical value.
*/
func main() {
	file, err := os.Open("099.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxVal := 0.0
	line, maxLine := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++
		numbers := strings.Split(scanner.Text(), `,`)
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])
		c := float64(b) * math.Log(float64(a))
		if c > maxVal {
			maxVal = c
			maxLine = line
		}
	}
	fmt.Println(maxLine)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
