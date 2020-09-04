package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
	name       string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
	name       string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
	name       string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
	fmt.Println(b.noise)
}
func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println(c.noise)
}
func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func AnimalFactory(name string, species string) Animal {
	if species == "cow" {
		return Cow{"grass", "walk", "moo", name}
	} else if species == "bird" {
		return Bird{"worms", "fly", "peep", name}
	} else if species == "snake" {
		return Snake{"mice", "slither", "hsss", name}
	} else {
		return nil
	}
}

func main() {
	var command string
	var choice string
	var animal string

	fmt.Println("Enter <Command> <Name> <species/action> or 'x' to exit")
	farm := make(map[string]Animal)

	for command != "x" {
		fmt.Print("> ")
		fmt.Scan(&command, &animal, &choice)
		strings.ToLower(command)
		strings.ToLower(choice)
		strings.ToLower(animal)

		if command != "x" {
			if command == "query" {
				if choice == "speak" {
					farm[animal].Speak()
				} else if choice == "eat" {
					farm[animal].Eat()
				} else if choice == "move" {
					farm[animal].Move()
				} else {
					fmt.Println("Invalid command")
				}
			} else if command == "newanimal" {
				a := AnimalFactory(animal, choice)
				if(a == nil){
					fmt.Println("Error adding animal")
				}else{
					farm[animal] = a
				}
			}
		}
	}
}
