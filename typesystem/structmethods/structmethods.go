package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	surname string
}

func (p person) getName() string {
	return p.name + " " + p.surname
}

func (p *person) setName(fullName string) {
	splitted := strings.Split(fullName, " ")

	p.name = splitted[0]
	p.surname = splitted[1]
}

func main() {

	jojoChar := person{"Ermes", "Costello"}
	fmt.Println(jojoChar.getName())

	jojoChar.setName("Jolyne Cujoh")
	fmt.Println(jojoChar.getName())
}
