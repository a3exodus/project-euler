package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

/*
Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
*/
func main() {
	file, err := os.Open("013.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := new(big.Int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num := new(big.Int)
		num, _ = num.SetString(scanner.Text(), 10)
		sum.Add(sum, num)
	}
	fmt.Println(sum.String()[:10])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
