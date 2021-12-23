package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

// interfaces have implicit implementation
func (p person) toString() string {
	return p.name + " " + p.surname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(prt printable) {
	fmt.Println(prt.toString())
}

func main() {
	var thing printable = person{"Jolyne", "Cujoh"}
	fmt.Println(thing.toString())
	print(thing)

	// polymorphism
	thing = product{"iPad", 7000.0}
	fmt.Println(thing.toString())
	print(thing)
}
