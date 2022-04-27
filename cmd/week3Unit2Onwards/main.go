package main

import (
	"fmt"
	"log"

	"github.com/Vyom-Yadav/go-lang-tut/pkg/interfaces"
)

func main() {
	fmt.Println(`
	Welcome, please enter first argument, name of the 
	animal (eg - "Brian", etc), and the third argument.
	`)
	fmt.Println(`
	Also one request, please don't use same name for animals,
	I have used a Map, now I clearly can use a slice but why to
	increase time complexity unnecessarily, you won't name your
	dog and cat pussy, would you?
	If you still use same names, then your previous animal will
	leave your farm, it will be sad.?`)
	animals := make(map[string]interfaces.Animal)
	for {
		fmt.Print(">")
		var firstArgument, nameOfAnimal, thirdArgument string
		_, err := fmt.Scan(&firstArgument, &nameOfAnimal, &thirdArgument)
		if err != nil {
			log.Fatal(err)
		}
		if firstArgument == "newanimal" {
			var newAnimal interfaces.Animal
			switch thirdArgument {
			case "cow":
				var cow *interfaces.Cow = interfaces.InitCow(nameOfAnimal)
				newAnimal = cow
			case "bird":
				var bird *interfaces.Bird = interfaces.InitBird(nameOfAnimal)
				newAnimal = bird
			case "snake":
				var snake *interfaces.Snake = interfaces.InitSnake(nameOfAnimal)
				newAnimal = snake
			default:
				log.Println("Which species is that? I thought there are only birds, cows and snakes, please try again!")
				continue
			}
			animal, present := animals[nameOfAnimal]
			if present {
				byByMsg := "%s the %s is leaving your farm (⌣̩̩́_⌣̩̩̀), wipe your tears and wish him good luck!\n"
				switch sh := animal.(type) {
				case *interfaces.Cow:
					log.Printf(byByMsg, sh.Name, "cow")
				case *interfaces.Bird:
					log.Printf(byByMsg, sh.Name, "bird")
				case *interfaces.Snake:
					log.Printf(byByMsg, sh.Name, "snake")
				}
			}
			animals[nameOfAnimal] = newAnimal
			fmt.Println("Created it!")
		} else if firstArgument == "query" {
			animal, present := animals[nameOfAnimal]
			if !present {
				log.Println("There is no animal by this name, try again!")
				continue
			}
			switch thirdArgument {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				log.Printf("Nah %s can't %s, please enter a valid action.\n", nameOfAnimal, thirdArgument)
			}
		} else {
			log.Printf("%s in not a valid first argument, please try again!", firstArgument)
		}
	}
}
