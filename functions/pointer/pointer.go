package main

import "fmt"

// Pure function
func inc1(n int) {
	n++
}

// Impure function
func inc2(n *int) {
	*n++
}

func main() {

	n := 1
	inc1(n) // passed by value
	fmt.Println(n)
	inc2(&n) // passed by reference
	fmt.Println(n)
}
