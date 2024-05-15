package main

import "fmt"

type day struct {
	dayOfWeek  int
	dayOfMonth int
	month      int
	year       int
}

var daysInMonths []int

/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/
func compute(n int) []day {

	daysInMonths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	allDays := []day{}
	today := day{1, 1, 1, 1900}
	for i := 1; i < n; i++ {
		allDays = append(allDays, today)
		tomorrow := nextDay(today)
		today = tomorrow
	}

	return allDays
}

func nextDay(d day) day {
	nextD := day{d.dayOfWeek + 1, d.dayOfMonth + 1, d.month, d.year}

	if nextD.dayOfWeek > 7 {
		nextD.dayOfWeek = 1
	}

	var daysInMonth int
	if d.month == 2 && isLeapYear(d.year) {
		daysInMonth = 29
	} else {
		daysInMonth = daysInMonths[d.month-1]
	}

	if nextD.dayOfMonth > daysInMonth {
		nextD.dayOfMonth = 1
		nextD.month++
		if nextD.month > 12 {
			nextD.month = 1
			nextD.year++
		}
	}

	return nextD
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func main() {
	count := 0
	allDays := compute(100_000)
	for _, today := range allDays {
		if 1901 <= today.year && today.year <= 2000 {
			if today.dayOfMonth == 1 && today.dayOfWeek == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}
