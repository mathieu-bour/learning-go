package main

import (
	"fmt"
)

// Instructions: https://go.dev/tour/flowcontrol/8

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Round", i, "z=", z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(144))
}
