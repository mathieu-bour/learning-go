package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	for i := 0; i < 100; i++ {
		ch <- i
	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
