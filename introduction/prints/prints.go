package main

import "fmt"

func main() {
	fmt.Print("Print -> Same")
	fmt.Print(" line.")

	fmt.Println(" New")
	fmt.Println("Line.")

	x := 3.141516 // How to print this?

	// #1
	fmt.Println("The value of x is", x)
	// #2
	fmt.Printf("The value of x is %.2f\n", x)
	// #3
	xstr := fmt.Sprint(x)
	fmt.Println("The value of x is " + xstr)

	a := 1
	b := 1.9999
	c := false
	d := "string"
	fmt.Printf("\n%d %f %.1f %t %s\n", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v\n", a, b, c, d)
}
