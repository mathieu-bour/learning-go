package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (err *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(*err))
}

func Sqrt2(x float64) (float64, *ErrNegativeSqrt) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, &err
	}

	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Round", i, "z=", z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
