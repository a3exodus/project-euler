package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

type triplet struct {
	a int64
	b int64
	c int64
}

/*
 */
func getPerimeters(limit int64) map[int64]int {
	perimeters := map[int64]int{}
	triplets := []triplet{}
	for n := int64(1); n < int64(math.Sqrt(float64(limit))); n++ {
		for m := n + 1; m < int64(math.Sqrt(float64(limit))); m++ {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			p := a + b + c
			if p < limit {
				for k := int64(1); k < limit; k++ {
					if p*k > limit {
						break
					}

					if a <= b {
						triplets = append(triplets, triplet{a: a * k, b: b * k, c: c * k})
					} else {
						triplets = append(triplets, triplet{a: b * k, b: a * k, c: c * k})
					}
				}
			}
		}
	}

	// remove duplicate triplets
	sort.Slice(triplets, func(i, j int) bool {
		if triplets[i].a != triplets[j].a {
			return triplets[i].a < triplets[j].a
		}
		return triplets[i].b < triplets[j].b
	})
	slices.Compact(triplets)

	for _, triplet := range triplets {
		perimeter := triplet.a + triplet.b + triplet.c
		perimeters[perimeter]++
	}
	return perimeters
}

func main() {
	perimeters := getPerimeters(1_500_000)
	count := 0
	for _, num := range perimeters {
		if num == 1 {
			count++
		}
	}
	fmt.Println(count)
}
