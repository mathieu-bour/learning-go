package main

import (
	"fmt"
	"time"
)

var count int

func inc() {
	for {
		count++
		time.Sleep(1)
	}
}

func dec() {
	for {
		count--
		time.Sleep(1)
	}
}

// In this example, two go routines are concurrently respectively incrementing (via inc) and decrementing (via dec) a
// shared counter. After 1 second, the counter is printed. We cannot predict the value of the counter at the end of
// the program.
func main() {
	go inc()
	go dec()
	time.Sleep(1000)
	fmt.Println(count)
}
