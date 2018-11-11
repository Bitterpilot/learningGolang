package main

import (
	"fmt"
	"math"
)

// Sqrt compute the square root of x using a loop
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("iteration: %d Result: %g\n", i, z)
	}
	return z
}

func main() {
	x := 2.0
	fmt.Printf("Standard lib: %g\n     My Sqrt: %g\n", math.Sqrt(x), Sqrt(x))
}
