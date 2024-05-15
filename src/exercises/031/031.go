package main

import (
	"fmt"
	"strconv"
)

var count int = 0

/*
In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:
1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).
It is possible to make £2 in the following way:
1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/
func main() {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	generatePerms(200, "", coins)
	fmt.Println(count)
}

func generatePerms(credit int, seq string, coins []int) {
	if credit < 0 {
		return
	}

	if credit == 0 {
		count++
	}

	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		newCoins := []int{}
		for _, c := range coins {
			if c <= coin {
				newCoins = append(newCoins, c)
			}
		}
		generatePerms(credit-coin, seq+strconv.Itoa(coin), newCoins)
	}
}
