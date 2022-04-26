package functions

import "fmt"

type MiInt int

func (mi MiInt) Double() int {
	return int(mi * 2)
}

type Point struct {
	x int
	y int
}

func (p *Point) PrintCoordinate() {
	fmt.Println("X: ", p.x)
	fmt.Println("Y: ", p.y)
}

func (p *Point) InitPoint(x, y int) {
	p.x = x    // same as *p.x = x
	(*p).y = y // Whereever we use pointer with '.', dereferencing is done automatically.
}
