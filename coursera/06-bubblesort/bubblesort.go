package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(integers []int) {
	for {
		sorted := true

		for i := 0; i < len(integers)-1; i++ {
			if integers[i] > integers[i+1] {
				Swap(integers, i)
				sorted = false
			}
		}

		// fmt.Printf("%v\n", integers)

		if sorted {
			break
		}
	}
}

func Swap(integers []int, index int) {
	if index < 0 || index >= len(integers)-1 {
		panic("swap: out of bounds")
	}

	integers[index], integers[index+1] = integers[index+1], integers[index]
}

func main() {
	var data []int

	fmt.Print("Enter a sequence of up to 10 integers, space-separated: ")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	elements := strings.Split(input.Text(), " ")
	for _, v := range elements {
		if i, err := strconv.Atoi(v); err == nil {
			data = append(data, i)
		}
	}

	BubbleSort(data)

	for _, v := range data {
		fmt.Printf("%d ", v)
	}
}
