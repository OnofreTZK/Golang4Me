package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	type dynamic interface{}

	var thing2 dynamic = "Obj"
	fmt.Println(thing2)

	thing2 = course{"Bicheiro Profissional"}
	fmt.Println(thing2)
}
