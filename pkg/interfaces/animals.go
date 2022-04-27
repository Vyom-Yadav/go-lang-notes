package interfaces

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Bird struct {
	Name string
}

func InitBird(name string) *Bird {
	bird := new(Bird)
	bird.Name = name
	return bird
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	Name string
}

func InitSnake(name string) *Snake {
	snake := new(Snake)
	snake.Name = name
	return snake
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

type Cow struct {
	Name string
}

func InitCow(name string) *Cow {
	cow := new(Cow)
	cow.Name = name
	return cow
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}
