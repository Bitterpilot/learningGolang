package main

import (
	"fmt"
	"math"
)

// Sqrt compute the square root of x using a loop
func Sqrt(x float64) float64 {
	const precision = 0.00000001
	z := float64(1)
	y := float64(0)
	for i := 0; i < 10; i++ {
		y = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("iteration: %d Result: %g and y= %g\n", i, z, y)
		// If equal within 8 decimal places
		if diff := math.Abs(z - y); diff < precision {break}
	}
	return z
}

func main() {
	x := 16.55445
	fmt.Printf("Standard lib: %g\n     My Sqrt: %g\n", math.Sqrt(x), Sqrt(x))
}
