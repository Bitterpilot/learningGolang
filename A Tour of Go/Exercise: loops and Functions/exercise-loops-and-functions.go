package main

import (
	"math"
	"fmt"
)

// Sqrt compute the square root of x using a loop
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("iteration:", i, " Result:", z)
	}
	return z
}

func main() {
	x := 2.0
	fmt.Println(math.Sqrt(x), Sqrt(x))
}
