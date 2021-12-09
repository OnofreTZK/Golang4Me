package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice is not an array, Slide defines a piece of an array
	s2 := a2[1:3] // slice start at index 1(inclusive) to index 3(exclusive)
	fmt.Println(a2, s2)

	s3 := a2[:2] // New slice pointing for the array 2
	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s2, s4)

	// Once arrays are static it is not worth instancing many if you need to change de values and
	// the size, so golang use slices to manipulate arrays
}
