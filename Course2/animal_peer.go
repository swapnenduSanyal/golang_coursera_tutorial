package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

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
	var choice string
	var animal string
	fmt.Println("Enter <Animal> <information> or 'x' to exit")

	farm := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for animal != "x" {
		fmt.Print("> ")
		fmt.Scan(&animal, &choice)
		strings.ToLower(choice)
		strings.ToLower(animal)
		if animal != "x" {
			if choice == "speak" {
				farm[animal].Speak()
			} else if choice == "eat" {
				farm[animal].Eat()
			} else if choice == "move" {
				farm[animal].Move()
			} else {
				fmt.Println("Invalid command")
			}
		}
	}
}
