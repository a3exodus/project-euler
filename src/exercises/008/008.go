package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
The four adjacent digits in the 1000-digit number that have the greatest product are 9x9x8x9 = 5832.
Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?
*/
func main() {
	file, err := os.Open("008.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	num := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	subLength := 13
	maxProduct := int64(0)
	for i := 0; i <= len(num)-subLength; i++ {
		subNum := num[i : i+subLength]
		prod := prodDigits(subNum)
		if prod > maxProduct {
			maxProduct = prod
		}
	}
	fmt.Print(maxProduct)
}

func prodDigits(num string) int64 {
	var prod int64 = 1
	for _, r := range num {
		prod *= int64(r - '0')
	}
	return prod
}
