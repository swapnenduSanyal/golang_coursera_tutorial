package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, sound string
}
type Snake struct {
	food, locomotion, sound string
}
type Bird struct {
	food, locomotion, sound string
}

func GetDetail(an Animal, details string) {
	switch details {
	case "eat":
		an.Eat()
	case "move":
		an.Move()
	case "speak":
		an.Speak()
	}
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}

func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func (b Bird) Speak() {
	fmt.Println(b.sound)
}

func MakeCow() Cow {
	var c Cow
	c.locomotion = "grass"
	c.food = "walk"
	c.sound = "moo"
	return c
}

func MakeBird() Bird {
	var b Bird
	b.locomotion = "fly"
	b.food = "worms"
	b.sound = "peep"
	return b
}

func MakeSnake() Snake {
	var s Snake
	s.locomotion = "slither"
	s.food = "mice"
	s.sound = "hsss"
	return s
}

func main() {
	var requestType, animalName, detailType string
	animals := make(map[string]Animal)
	for {
		fmt.Print("> ")
		fmt.Scan(&requestType, &animalName, &detailType)
		if requestType == "newanimal"{
			switch detailType{
			case "cow": animals[animalName] = MakeCow()
			case "bird": animals[animalName] = MakeBird()
			case "snake":animals[animalName] = MakeSnake()
			}
			fmt.Println("Created it!")
		} else if requestType == "query"{
			switch an := animals[animalName].(type) {
			case Cow: GetDetail(an,detailType)
			case Bird: GetDetail(an,detailType)
			case Snake: GetDetail(an,detailType)
			}
		}

	}
}
