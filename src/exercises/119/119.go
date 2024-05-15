package main

import (
	"fmt"
	"math/big"
	"sort"
)

/*
The number 512 is interesting because it is equal to the sum of its digits raised to some power: 5 + 1 + 2 = 8, and 8^3=512.
Another example of a number with this property is 614656=28^4.

We shall define a_n to be the nth term of this sequence and insist that a number must contain at least two digits to have a sum.
You are given that a_2=512 and a_10=61465. Find a_30.
*/
func main() {

	var sequence []*big.Int
	pow := new(big.Int)
	for a := 2; a <= 100; a++ {
		bigA := big.NewInt(int64(a))
		for b := 2; b <= 100; b++ {
			bigB := big.NewInt(int64(b))
			pow.Exp(bigA, bigB, nil)
			sum := digitSum(pow)
			if a == sum {
				sequence = append(sequence, new(big.Int).Set(pow))
			}
		}
	}
	sort.Slice(sequence, func(a, b int) bool {
		return sequence[a].Cmp(sequence[b]) < 0
	})
	fmt.Println(sequence[29])
}

func digitSum(n *big.Int) int {
	sum := 0
	for _, r := range n.String() {
		sum += int(r - '0')
	}
	return sum
}
