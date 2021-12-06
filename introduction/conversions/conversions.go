package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// Int -> String
	fmt.Println("Testing string with int -> " + string(20)) // Converts to ascii not to integer
	// Use this
	fmt.Printf("Testing string with int -> %d\n", 20)
	// Or this
	fmt.Println("Testing string with int -> " + strconv.Itoa(20))

	// String -> Integer
	num, _ := strconv.Atoi("20") // Atoi return a "Result.t" like in OCaml -> (int, error)
	// The anon var _ is to ignore the error

	fmt.Println(num - 19)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("True")
	}

}
