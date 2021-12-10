package main

import "fmt"

var sum = func(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(sum(75, 25))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))

}
