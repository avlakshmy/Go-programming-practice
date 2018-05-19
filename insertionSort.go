//This program implements insertion sort in Go

package main

import "fmt"

var arr = []int{12, 6, 78, 66, 2, 90, 45, 55, 3, 83}

func insertionSort(array []int) []int{
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i-1
		
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j--
		}
		
		array[j+1] = temp
	}
	
	return array		
}

func main() {
	fmt.Println(insertionSort(arr))
}
