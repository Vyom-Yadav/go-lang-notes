package functions

import (
	"fmt"
)

type Animals struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animals) InitAnimal(animalName string) {
	switch animalName {
	case "cow":
		animal.food = "grass"
		animal.locomotion = "walk"
		animal.noise = "moo"
	case "bird":
		animal.food = "worms"
		animal.locomotion = "fly"
		animal.noise = "peep"
	case "snake":
		animal.food = "mice"
		animal.locomotion = "slither"
		animal.noise = "hsss"
	}
}

func (animal *Animals) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animals) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animals) Speak() {
	fmt.Println(animal.noise)
}
