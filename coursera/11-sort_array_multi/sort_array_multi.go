package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SortPartition(input []int, c chan []int) {
	fmt.Printf("I am sorting %v\n", input)
	sort.Ints(input)
	c <- input
}

func main() {
	fmt.Println("Enter as any integers as you want, space separated:")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		panic(err)
	}

	var integers []int

	for _, v := range strings.Split(input, " ") {
		conv, err := strconv.Atoi(strings.TrimSpace(v))

		if err != nil {
			fmt.Println(err)
			continue
		}

		integers = append(integers, conv)
	}

	fmt.Printf("Sorting %v\n", integers)

	var ints0, ints1, ints2, ints3 []int

	for i, v := range integers {
		switch i % 4 {
		case 0:
			ints0 = append(ints0, v)
		case 1:
			ints1 = append(ints1, v)
		case 2:
			ints2 = append(ints2, v)
		case 3:
			ints3 = append(ints3, v)
		}
	}

	c := make(chan []int)
	go SortPartition(ints0, c)
	go SortPartition(ints1, c)
	go SortPartition(ints2, c)
	go SortPartition(ints3, c)

	merged := append(<-c, <-c...)
	merged = append(merged, <-c...)
	merged = append(merged, <-c...)

	sort.Ints(merged)

	fmt.Printf("Sorting result: %v\n", merged)
}
