package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

var cow = Animal{"grass", "walk", "moo"}
var bird = Animal{"worms", "fly", "peep"}
var snake = Animal{"mice", "slither", "hsss"}
var mappings = map[string]Animal{"cow": cow, "bird": bird, "snake": snake}

func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var name, action string

	for {
		fmt.Println("What do you want to learn?")
		fmt.Println("Enter an animal (cow|bird|snake) and an action (eat|move|noise, space separated)")
		fmt.Print("> ")

		if _, err := fmt.Scan(&name, &action); err != nil {
			panic(err)
		}

		if name != "cow" && name != "bird" && name != "snake" {
			fmt.Println("Invalid animal, please try again")
			continue
		}

		switch action {
		case "eat":
			mappings[name].Eat()
		case "move":
			mappings[name].Move()
		case "noise":
			mappings[name].Speak()
		default:
			fmt.Println("Invalid action, please try again")
		}
	}
}
