package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (_ Cow) Eat() {
	fmt.Println("grass")
}
func (_ Cow) Move() {
	fmt.Println("walk")
}
func (_ Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (_ Bird) Eat() {
	fmt.Println("worms")
}
func (_ Bird) Move() {
	fmt.Println("fly")
}
func (_ Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (_ Snake) Eat() {
	fmt.Println("mice")
}
func (_ Snake) Move() {
	fmt.Println("slither")
}
func (_ Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := map[string]Animal{}

	for {
		fmt.Print("> ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		args := strings.Split(strings.TrimSpace(input.Text()), " ")

		switch args[0] {
		case "newanimal":
			if len(args) != 3 {
				fmt.Printf("Invalid args count, expected 3 got %d\n", len(args))
				continue
			}

			// args[2] is the animal type
			switch args[2] {
			case "cow":
				animals[args[1]] = Cow{}
			case "bird":
				animals[args[1]] = Bird{}
			case "snake":
				animals[args[1]] = Snake{}
			default:
				fmt.Printf("Invalid animal type %s\n", args[1])
				continue
			}

			fmt.Println("Created it!")

		case "query":
			if len(args) != 3 {
				fmt.Printf("Invalid args count, expected 3 got %d\n", len(args))
				continue
			}

			found, ok := animals[args[1]]

			if !ok {
				fmt.Printf("cannot find animal named %s\n", args[1])
				continue
			}

			switch args[2] {
			case "eat":
				found.Eat()
			case "move":
				found.Move()
			case "speak":
				found.Speak()
			default:
				fmt.Printf("Invalid action %s\n", args[2])
				continue
			}

		default:
			fmt.Println("Unexpected command")
		}
	}
}
