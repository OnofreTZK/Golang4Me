package main

import (
	"fmt"
)

func fact(n uint) (uint, error) {

	// Assigning the closure
	var aux func(uint, uint) (uint, error)

	aux = func(num, total uint) (uint, error) {

		switch {
		case num == 0:
			return 1, nil
		case num == 1:
			return total, nil
		case num == n:
			return aux(num-1, num*total)
		default:
			return aux(num-1, (num-1)*total)
		}
	}

	return aux(n, n-1)
}

func main() {
	val, _ := fact(10)
	fmt.Println(val)
	val2, _ := fact(0)
	fmt.Println(val2)
}
