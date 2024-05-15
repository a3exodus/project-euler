package main

import "fmt"

/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be
1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24.
By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers.
However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.
Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/
func main() {
	n := 28123

	abundants := []int{}
	candidates := map[int]bool{}

	for i := 1; i <= n; i++ {
		candidates[i] = true
		if d(i) > i {
			abundants = append(abundants, i)
		}
	}

	for _, a := range abundants {
		for _, b := range abundants {
			sum := a + b
			if sum <= n {
				candidates[sum] = false
			}
		}
	}

	sum := 0
	for num, v := range candidates {
		if v {
			sum += num
		}
	}
	fmt.Println(sum)
}

func d(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
