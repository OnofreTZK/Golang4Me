package main

import "fmt"

type sport interface {
	enableTurbo()
}

type luxury interface {
	enableAutoPilot()
}

type luxurysport interface {
	sport
	luxury
	// Implicited inherited methods signatures
	// we can add more methods too
}

type gle63 struct{}

func (car gle63) enableTurbo() {
	fmt.Println("Fuuuuiin")
}

func (car gle63) enableAutoPilot() {
	fmt.Println("I'm bout to crash this car...")
}

func main() {
	var car luxurysport = gle63{}
	car.enableTurbo()
	car.enableAutoPilot()
}
