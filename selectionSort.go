//This program implements selection sort in Go

package main

import "fmt"

var arr = []int{12, 6, 78, 66, 2, 90, 45, 55, 3, 83}

func selectionSort(array []int) []int{
	for i := 0; i < len(array); i++ {
		k := retMinIndex(array, i)
		array[i], array[k] = array[k], array[i]
	}
	return array		
}

func retMinIndex (array []int, ind int) int {
	min := ind
	for i := ind+1; i < len(array); i++ {
		if array[min] > array[i] {
			min = i
		}
	}
	return min
}

func main() {
	fmt.Println(selectionSort(arr))
}
