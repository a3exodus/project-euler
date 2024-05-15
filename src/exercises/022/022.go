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
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
So, COLIN would obtain a score of 983 x 53 = 49714.

What is the total of all the name scores in the file?
*/
func main() {
	file, err := os.Open("022.txt")
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

	sum := int64(0)
	for index, word := range words {
		sum += int64(index+1) * int64(wordValue(word))
	}
	fmt.Println(sum)
}

func wordValue(word string) int {
	sum := 0
	for _, r := range word {
		sum += int(r - 'A' + 1)
	}
	return sum
}
