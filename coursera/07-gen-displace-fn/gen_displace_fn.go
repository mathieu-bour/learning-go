package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, initialSpeed, initialDisplacement float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*math.Pow(t, 2) + initialSpeed*t + initialDisplacement
	}
}

func main() {
	var acc, initialSpeed, initialDisplacement, time float64
	fmt.Print("Enter values for acceleration, initial velocity, and initial displacement (space seperated): ")

	_, err := fmt.Scan(&acc, &initialSpeed, &initialDisplacement)
	if err != nil {
		panic(err)
	}

	fn := GenDisplaceFn(acc, initialSpeed, initialDisplacement)

	fmt.Print("Enter time t=")
	_, err = fmt.Scan(&time)

	if err != nil {
		panic(err)
	}

	fmt.Println(fn(time))
}
