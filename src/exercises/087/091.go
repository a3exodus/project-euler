package main

import "fmt"

type vector struct {
	x int
	y int
}

/*
The points P(x1,y1) and Q(x2,y2) are plotted at integer co-ordinates and are joined to the origin, O(0,0) to form OPQ.
There are exactly fourteen triangles containing a right angle that can be formed when each co-ordinate lies between 0 and 2 inclusive; that is, 0 <= x1,x2,y1,y2 <= 2.
Given that 0 <= x1,x2,y1,y2 <= 50, how many right triangles can be formed?
.
*/
func compute(n int) int {
	coordinates := []vector{}
	for x := 0; x <= n; x++ {
		for y := 0; y <= n; y++ {
			coordinates = append(coordinates, vector{x, y})
		}
	}
	coordinates = coordinates[1:]

	count := 0
	for _, x := range coordinates {
		for _, y := range coordinates {
			if x == y {
				continue
			}

			if dotProduct(x, y) == 0 {
				count++
				continue
			}

			z := minusVector(x, y)
			if dotProduct(x, z) == 0 || dotProduct(y, z) == 0 {
				count++
			}
		}
	}

	return count / 2
}

func minusVector(x vector, y vector) vector {
	return vector{x.x - y.x, x.y - y.y}
}

func dotProduct(x vector, y vector) int {
	return x.x*y.x + x.y*y.y
}

func main() {
	fmt.Println(compute(50))
}
