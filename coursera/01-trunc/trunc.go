package main

import "fmt"

func main() {
	input := new(float64)
	_, err := fmt.Scan(input)
	if err != nil {
		panic("Invalid float")
	}

	integer := int(*input)
	fmt.Printf("%d", integer)
}
