package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Each character on a computer is assigned a unique code and the preferred standard is ASCII (American Standard Code for Information Interchange).
For example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.
A modern encryption method is to take a text file, convert the bytes to ASCII, then XOR each byte with a given value, taken from a secret key.
The advantage with the XOR function is that using the same encryption key on the cipher text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.
For unbreakable encryption, the key is the same length as the plain text message, and the key is made up of random bytes.
The user would keep the encrypted message and the encryption key in different locations, and without both "halves", it is impossible to decrypt the message.
Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key.
If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the message.
The balance for this method is using a sufficiently long password key for security, but short enough to be memorable.
Your task has been made easy, as the encryption key consists of three lower case characters.
Using 0059_cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes, and the knowledge that the plain text must contain common English words, decrypt the message and find the sum of the ASCII values in the original text.
*/
func main() {
	file, err := os.Open("059.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	var words []string
	cipherText := []int{}
	for scanner.Scan() {
		line = scanner.Text()
		words = strings.Split(line, `,`)
		for _, word := range words {
			cipherByte, _ := strconv.Atoi(word)
			cipherText = append(cipherText, int(cipherByte))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	maxScore := 0
	key := []int{}
	for a := 97; a <= 122; a++ {
		for b := 97; b <= 122; b++ {
			for c := 97; c <= 122; c++ {
				potentialKey := []int{a, b, c}
				plainText := decrypt(cipherText, potentialKey)
				score := cryptanalysisScore(plainText)
				if score > maxScore {
					maxScore = score
					key = potentialKey
				}
			}
		}
	}

	sum := 0
	plainText := decrypt(cipherText, key)
	for _, byte := range plainText {
		sum += byte
	}
	fmt.Println(sum)
}

func decrypt(cipherText []int, key []int) []int {
	plainText := []int{}
	for i := 0; i < len(cipherText); i++ {
		keyByte := key[i%3]
		plainText = append(plainText, cipherText[i]^keyByte)
	}
	return plainText
}

func cryptanalysisScore(plainText []int) int {
	count := 0
	for _, byte := range plainText {
		if 65 <= byte && byte <= 122 {
			count++
		}
	}
	return count
}
