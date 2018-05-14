//This program takes two DNA strands' base sequences, and determines their Hamming distance
//i.e., the number of positions where they differ

package main

import "fmt"

func hammingDistance(str1, str2 string) int {
	var d int
	d = 0
	
	for i := 0; i < len(str1); i++ {
		c1 := rune([]rune(str1)[i])
		c2 := rune([]rune(str2)[i])

		if c1 != c2 {
			d++
		}
	}
	return d	
}

func main() {
	var str1 string
	str1 = "TAGCCGTACGTACGTTGCA"
	var str2 string
	str2 = "TAGCCCTACGCACGCTGGA"

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(hammingDistance(str1, str2))
}
