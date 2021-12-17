package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	car          // Anon attributes -> every attribute of car is visible for ferrari
	turbaoVeneno bool
}

func main() {
	f := ferrari{}
	f.name = "Italia"
	f.currentSpeed = 0
	f.turbaoVeneno = true
	fmt.Println(f)
}
