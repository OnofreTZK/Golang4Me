package main

import "fmt"

func swap(p1, p2 int) (snd, fst int) {
	snd = p2
	fst = p1

	return // clean return
}

func main() {

	v1, v2 := 5, 10
	fmt.Println(v1, v2)
	v3, v4 := swap(v1, v2)
	fmt.Println(v3, v4)

}
