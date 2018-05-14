//This program takes a year and prints out if it is a leap year or not

package main

import "fmt"

func isLeapYear (year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func main() {
	years := [4]int{1996, 1993, 1900, 2000}
	for i := 0; i < 4; i++ {
		fmt.Println(years[i], isLeapYear(years[i]))
	}
}
	
