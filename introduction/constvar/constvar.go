package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // Inferred by the compiler as a float

	// reduced form to declare a variable in go
	area := PI * math.Pow(raio, 2)
	fmt.Printf("Circle area: %f\n", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "string"

	fmt.Println(g, h, i)
}
