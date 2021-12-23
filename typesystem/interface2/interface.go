package main

import "fmt"

type xunning interface {
	enableTurbo()
}

type volkswagem struct {
	model     string
	isTurboOn bool
	speed     int
}

func (car *volkswagem) enableTurbo() {
	car.isTurboOn = true
}

func main() {
	car := volkswagem{"Gol GTI Quadrado", false, 0}
	car.enableTurbo()

	// Polymorphism has to be explicit -> change the reference
	//var car2 volkswagem = &volkswagem{"Fuscao Envenenado", false, 0}
	// Nowadays it can be!
	var car2 volkswagem = volkswagem{"Fuscao Envenenado", false, 0}
	car2.enableTurbo()

	fmt.Println(car, car2)
}
