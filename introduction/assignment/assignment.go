package main

import "fmt"

func main() {

	var b byte = 3 // Explicit typed assignment
	fmt.Println(b)

	i := 3 // Implicit typed assignment
	i += 3 // Increment
	i -= 3 // Decrement
	i /= 2 // Divide
	i *= 2 // Multiply
	i %= 2 // Remain of the division

	x, y := 1, 2 // Multi assignment
	fmt.Println(x, y)

	x, y = y, x // Swap the values between the 2 vars
	fmt.Println(x, y)
}
