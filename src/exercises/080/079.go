package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
A common security method used for online banking is to ask the user for three random characters from a passcode.
For example, if the passcode was 531278, they may ask for the 2nd, 3rd, and 5th characters; the expected reply would be: 317.
The text file, keylog.txt, contains fifty successful login attempts.
Given that the three characters are always asked for in order, analyse the file so as to determine the shortest possible secret passcode of unknown length.
*/
func main() {
	file, err := os.Open("079.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passwords := []int{}
	scanner := bufio.NewScanner(file)
	count := 0
	var line int
	for scanner.Scan() {
		line, _ = strconv.Atoi(scanner.Text())
		passwords = append(passwords, line)
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var foundPass bool
	for i := 0; i <= 100_000_000; i++ {
		foundPass = true
		for _, password := range passwords {
			if !isPotentialPassword(strconv.Itoa(i), strconv.Itoa(password)) {
				foundPass = false
				break
			}
		}
		if foundPass {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}

func isPotentialPassword(candidate string, key string) bool {
	hit := 0
	rKey := []rune(key)[hit]
	for _, r := range candidate {
		if r == rKey {
			hit++
			if hit == len(key) {
				return true
			}
			rKey = []rune(key)[hit]
		}
	}
	return false
}
