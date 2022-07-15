package main

import "fmt"

// See https://go.dev/tour/moretypes/26

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var i, j = 0, 1
	k := -1

	return func() int {
		k++

		switch k {
		case 0:
			return 0
		case 1:
			return 1
		default:
			i, j = j, i+j
			return j
		}
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
