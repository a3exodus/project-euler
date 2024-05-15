package main

import "fmt"

func main() {
	maxSols, maxP := 0, 0
	for p := 3; p <= 1000; p++ {
		sols := numSols(p)
		if sols > maxSols {
			maxSols = sols
			maxP = p
		}
	}
	fmt.Println(maxP)
}

func numSols(p int) int {
	sols := 0
	for a := 1; a <= p/3; a++ {
		for b := a + 1; b <= p/2; b++ {
			c := p - b - a
			if a*a+b*b == c*c {
				sols++
			}
		}
	}
	return sols
}
