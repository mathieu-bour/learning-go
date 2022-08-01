package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	people := make([]Person, 0)
	var filename string
	fmt.Print("Enter the filename to open: ")
	_, err := fmt.Scan(&filename)

	if err != nil {
		fmt.Printf("unable to scan %v", err)
		return
	}

	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("unable to open file %s, %v", filename, err)
		return
	}

	fscanner := bufio.NewScanner(f)
	fscanner.Split(bufio.ScanLines)

	for fscanner.Scan() {
		line := strings.TrimSpace(fscanner.Text())
		parts := strings.Split(line, " ")
		people = append(people, Person{fname: parts[0], lname: parts[1]})
	}

	err = f.Close()

	if err != nil {
		fmt.Printf("unable to close file %s: %v", filename, err)
		return
	}

	for i, p := range people {
		fmt.Printf("%d - first name: %s - last name: %s\n", i, p.fname, p.lname)
	}
}
