package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := map[string]string{}

	fmt.Print("Enter your first name: ")
	firstname, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	data["firstname"] = strings.TrimSpace(firstname)

	fmt.Print("Enter your address: ")
	address, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	data["address"] = strings.TrimSpace(address)

	value, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Output JSON: %s", string(value))
}
