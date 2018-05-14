//This program implements the square root function using Newton's method

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	zp := 1.0
	zn := 0.0
	for math.Abs(zn - zp) > 0.0000000000001 {
		zp = z
		z -= (z*z - x) / (2*z)
		zn = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(8))
}
