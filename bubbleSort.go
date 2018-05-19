//This program implements bubble sort in Go

package main

import "fmt"

var arr = []int{12, 6, 78, 66, 2, 90, 45, 55, 3, 83}

func bubbleSort(array []int) []int{
	for i := 0; i < len(array) - 1; i++ {
		for j := i+1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array		
}

func main() {
	fmt.Println(bubbleSort(arr))
}
