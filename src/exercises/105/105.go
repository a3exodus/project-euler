package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var perms []string

type subsetSum struct {
	numElements int
	sum         int
}

/*
Let S(A) represent the sum of elements in set A of size n.
We shall call it a special sum set if for any two non-empty disjoint subsets, B and C, the following properties are true:

1. S(B) ≠ S(C); that is, sums of subsets cannot be equal.
2. If B contains more elements than C then S(B) > S(C).

For example, {81, 88, 75, 42, 87, 84, 86, 65} is not a special sum set because 65 + 87 + 88 = 75 + 81 + 84, whereas {157, 150, 164, 119, 79, 159, 161, 139, 158} satisfies both rules for all possible subset pair combinations and S(A) = 1286.
Using sets.txt (right click and "Save Link/Target As..."), a 4K text file with one-hundred sets containing seven to twelve elements (the two examples given above are the first two sets in the file), identify all the special sum sets, A1, A2, ..., Ak, and find the value of S(A1) + S(A2) + ... + S(Ak).
NOTE: This problem is related to Problem 103 and Problem 106.
*/
func main() {
	file, err := os.Open("105.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		A := []int{}
		numbers := strings.Split(line, `,`)
		for _, num := range numbers {
			number, _ := strconv.Atoi(num)
			A = append(A, number)
		}
		if isSpecialSumSet(A) {
			subsetSums := getSubsetSums(A)
			sum := subsetSums[len(subsetSums)-1].sum
			totalSum += sum
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalSum)
}

func isSpecialSumSet(A []int) bool {
	n := len(A)
	perms = []string{}
	generatePerms("", n)

	subsetSums := getSubsetSums(A)

	// n+1 because (0,0) is also a subsetsum
	for b := 0; b < len(subsetSums)-1; b++ {
		c := b + 1
		subsetSumB := subsetSums[b]
		subsetSumC := subsetSums[c]

		if subsetSumB == subsetSumC {
			return false
		}
		if subsetSumB.numElements < subsetSumC.numElements && subsetSumB.sum >= subsetSumC.sum {
			return false
		}
	}

	return true
}

func getSubsetSums(A []int) []subsetSum {
	var subsetSums []subsetSum
	for _, perm := range perms {
		subsetSums = append(subsetSums, getSubsetSum(A, perm))
	}

	sort.Slice(subsetSums, func(b, c int) bool {
		numElementsB := subsetSums[b].numElements
		numElementsC := subsetSums[c].numElements
		if numElementsB != numElementsC {
			return numElementsB < numElementsC
		}
		return subsetSums[b].sum < subsetSums[c].sum
	})

	return subsetSums
}

func getSubsetSum(A []int, perm string) subsetSum {
	n := len(A)
	count, sum := 0, 0
	for i := 0; i < n; i++ {
		useNum, _ := strconv.Atoi(perm[i : i+1])
		count += useNum
		sum += useNum * A[i]
	}
	return subsetSum{count, sum}
}

func generatePerms(seq string, n int) {
	l := len(seq)
	if l == n {
		perms = append(perms, seq)
		return
	}

	generatePerms(seq+"0", n)
	generatePerms(seq+"1", n)
}
