//This program implements conversion from decimal to binary and vice-versa

package main

import "fmt"
import "strconv"

func DecToBin (n int) int {
	var s string
	var d int
	for n > 0 {
		n, d = n/2, n%2
		s1 := strconv.Itoa(d)
		s = s1 + s
	}
	
	bin, _ := strconv.Atoi(s)
	return bin
} 

func BinToDec (b uint) int {
	var i uint
	var d uint
	var s uint
	i = 0
	s = 0
	for b > 0 {
		d = b%10
		s = s + (d << i)
		i++
		b = b/10
	}
	return int(s)
}



func main() {
	fmt.Println(DecToBin(25))
	fmt.Println(BinToDec(11001))
}

