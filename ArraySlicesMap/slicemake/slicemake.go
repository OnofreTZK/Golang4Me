package main

import "fmt"

func main() {
	s := make([]int, 10) // A slice instanced without point to an array points to a
	//internal array created by the compiler
	s[9] = 12 // slices accept assignment by indexing
	fmt.Println(s)

	s = make([]int, 10, 20) // Slice with 10 elements pointing to a internal array with 20 elements
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// Appending after reached the max cap
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) // Behaviour like <vector> from C++
}
