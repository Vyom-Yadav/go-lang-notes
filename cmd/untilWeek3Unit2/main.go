package main

import (
	"fmt"
	"log"

	"github.com/Vyom-Yadav/go-lang-tut/pkg/functions"
)

func main() {
	fmt.Println("No function called by default, you can change it with the function you wish to call!")
}

func inputAnimals() {
	fmt.Print(">")
	var animalName, animalTask string
	fmt.Scan(&animalName, &animalTask)
	var cow, bird, snake *functions.Animals
	var currentAnimal **functions.Animals
	var isCowInit, isBirdInit, isSnakeInit bool
	switch animalName {
	case "cow":
		cow = initAnimal(animalName, &currentAnimal, &isCowInit, cow)
	case "bird":
		bird = initAnimal(animalName, &currentAnimal, &isBirdInit, bird)
	case "snake":
		snake = initAnimal(animalName, &currentAnimal, &isSnakeInit, snake)
	default:
		log.Fatal("Exiting, invalid animal")
	}
	for {
		switch animalTask {
		case "eat":
			(*currentAnimal).Eat()
		case "move":
			(*currentAnimal).Move()
		case "speak":
			(*currentAnimal).Speak()
		default:
			log.Fatal("Exiting, invalid task")
		}
		fmt.Printf("%p", *currentAnimal)
		fmt.Print(">")
		fmt.Scan(&animalName, &animalTask)

		switch animalName {
		case "cow":
			if !isCowInit {
				fmt.Println("Here")
				cow = initAnimal(animalName, &currentAnimal, &isCowInit, cow)
			}
			currentAnimal = &cow
		case "bird":
			if !isBirdInit {
				bird = initAnimal(animalName, &currentAnimal, &isBirdInit, bird)
			}
			currentAnimal = &bird
		case "snake":
			if !isSnakeInit {
				snake = initAnimal(animalName, &currentAnimal, &isSnakeInit, snake)
			}
			currentAnimal = &snake
		default:
			log.Fatal("Exiting, invalid animal")
		}

	}
}
func initAnimal(name string, currentAnimalPointerUpdater ***functions.Animals, animalSetFlag *bool, currAnimal *functions.Animals) *functions.Animals {
	var animal functions.Animals
	animal.InitAnimal(name)
	*animalSetFlag = true
	*currentAnimalPointerUpdater = &currAnimal
	(*(*currentAnimalPointerUpdater)) = &animal
	return &animal
}

func myPointTest() {
	var point functions.Point
	point.InitPoint(1, 3) // We need not explicitly do (&point).InitPoint(1,3), go does that for us, it is done implicitly.
	point.PrintCoordinate()
}

func myIntTest() {
	miInt := functions.MiInt(3)
	fmt.Println(miInt)
	asd := miInt.Double() // miInt is implicitly passed by value, whatever on left side of dot is implicitly passed by value
	fmt.Println(miInt)
	fmt.Println(asd)
}

func inputForBubble() {
	sli := make([]int, 0, 10)
	fmt.Println("Please Enter the numbers: ")
	for i := 0; i < 10; i++ {
		var val int
		fmt.Scan(&val)
		sli = append(sli, val)
	}
	functions.BubbleSort(sli)
	fmt.Println(sli)
}

func inputForDisp() {
	var acc, initVel, initDisp, time float64
	fmt.Println("Enter Acceleration, initial velocity and initial displacement: ")
	fmt.Scan(&acc, &initVel, &initDisp)
	fn := functions.GenDispFn(acc, initVel, initDisp)
	fmt.Println("Enter the value of time: ")
	fmt.Scan(&time)
	fmt.Println(fn(time))
}
