package main

import "fmt"

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which, a^2+b^2=c^2.
There exists exactly one Pythagorean triplet for which a + b + c = 1000. Find the product abc.
*/
func main() {
outer:
	for a := 1; a < 333; a++ {
		for b := a + 1; b < 500; b++ {
			c := 1000 - b - a
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				break outer
			}
		}
	}
}
