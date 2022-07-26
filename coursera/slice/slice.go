package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	data := make([]int, 0, 3)

	for {
		input := new(string)
		fmt.Print("Enter a integer (or \"x\" to exit): ")
		_, err := fmt.Scan(input)

		if err != nil {
			panic(err)
		}

		if *input == "x" {
			break
		}

		integer, err := strconv.Atoi(*input)

		if err != nil {
			panic(err)
		}

		data = append(data, integer)

		sort.Ints(data) // Sort the slice
		fmt.Printf("Output slice: %#v\n", data)
	}
}
