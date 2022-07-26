package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter test string: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		panic(err)
	}

	// test is in lowercase
	test := strings.TrimSpace(strings.ToLower(input))

	if strings.HasPrefix(test, "i") && strings.Contains(test, "a") && strings.HasSuffix(test, "n") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
