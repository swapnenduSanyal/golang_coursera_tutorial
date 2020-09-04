package main

import "fmt"

type animal struct {
	food, locomotion, noise string
}

func (an animal) GetDetail(details string){
	switch details {
	case "eat": an.Eat()
	case "move": an.Move()
	case "speak": an.Speak()
	}
}

func (an animal) Eat(){
	fmt.Println(an.food)
}

func (an animal) Move(){
	fmt.Println(an.locomotion)
}

func (an animal) Speak(){
	fmt.Println(an.noise)
}

func main() {
	cow := animal{
		locomotion: "grass",
		food:       "walk",
		noise:      "moo"}
	bird := animal{
		locomotion: "fly",
		food:       "worms",
		noise:      "peep"}
	snake := animal{
		locomotion: "slither",
		food:       "mice",
		noise:      "hsss"}
	var animalType, detailType string
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s ", &animalType, &detailType)
		switch animalType {
		case "cow":
			cow.GetDetail(detailType)
		case "snake":
			snake.GetDetail(detailType)
		case "bird":
			bird.GetDetail(detailType)
		}
	}
}
