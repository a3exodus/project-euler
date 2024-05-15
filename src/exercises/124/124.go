package main

import (
	"fmt"
	"project-euler/helpers"
	"sort"
)

type OrderedRad struct {
	index int
	value int
}

/*
The radical of n, rad(n), is the product of the distinct prime factors of n.
For example, 504=2^3x3^2x7, so rad(504)=2x3x7=42.
Let E(k) be the 𝑘 k-th element in the sorted n column of radicals; for example, E(4) = 8 and E(6) = 9.
If rad(n) is sorted for 1 ≤ n ≤ 100000, find E(10000).
*/
func compute(n int) []OrderedRad {
	orderedRads := make([]OrderedRad, n)
	for i := 1; i <= n; i++ {
		orderedRads[i-1] = OrderedRad{
			i,
			rad(i),
		}
	}
	sort.Slice(orderedRads, func(i, j int) bool {
		valI := orderedRads[i].value
		valJ := orderedRads[j].value
		if valI != valJ {
			return valI < valJ
		}
		return orderedRads[i].index < orderedRads[j].index
	})
	return orderedRads
}

func rad(n int) int {
	rad := 1
	for prime, _ := range helpers.Factorize(int64(n)) {
		rad *= int(prime)
	}
	return rad
}

func main() {
	OrderedRad := compute(100_000)
	fmt.Println(OrderedRad[10_000-1].index)
}
