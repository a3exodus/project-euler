package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t_(10).
If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/
func main() {
	file, err := os.Open("042.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	var words []string
	for scanner.Scan() {
		line = scanner.Text()
		words = strings.Split(line, `,`)
		for index, word := range words {
			words[index] = strings.ReplaceAll(word, "\"", "")
		}
		sort.Strings(words)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	triangles := map[int]bool{}
	for i := 1; i < 1000; i++ {
		triangle := i * (i + 1) / 2
		triangles[triangle] = true
	}

	count := 0
	for _, word := range words {
		val := wordValue(word)
		_, exists := triangles[val]
		if exists {
			count++
		}
	}
	fmt.Println(count)
}

func wordValue(word string) int {
	sum := 0
	for _, r := range word {
		sum += int(r - 'A' + 1)
	}
	return sum
}
