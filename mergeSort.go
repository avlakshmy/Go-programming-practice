//This program implements merge sort in Go

package main

import "fmt"

var arr = []int{12, 6, 78, 66, 2, 90, 45, 55, 3, 83}

func mergesort(array []int, start int, end int) []int {
	if start < end {		
		var mid int 
		mid = (start + end)/2;		
		arr1 := mergesort(array, start, mid)
		arr2 := mergesort(array, mid+1, end)		
		return merge(arr1, arr2)
	} else if start == len(array){
		var ar = []int{array[start-1]}
		return ar
	} else{
		var ar = []int{array[start]}
		return ar
	}
}

func merge(arr1 []int, arr2 []int) []int {
	n1 := len(arr1)
	n2 := len(arr2)
	
	array := make([]int, n1+n2)
	var a, b, c int
	a, b, c = 0, 0, 0
	
	for a < n1 && b < n2 {
		if arr1[a] < arr2[b] {
			array[c] = arr1[a]
			a++
		} else {
			array[c] = arr2[b]
			b++
		}
		c++
	}
	
	for a < n1 {	
		array[c] = arr1[a]
		a++
		c++
	}
	
	for b < n2 {	
		array[c] = arr2[b]
		b++
		c++
	}	
	
	return array	
}	
			
func main() {
	fmt.Println(mergesort(arr, 0, len(arr)))
}
