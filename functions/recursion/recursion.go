package main

import (
	"fmt"
)

func fact(n int) (int, error) {

	var aux func(int, int) (int, error)

	aux = func(prev, total int) (int, error) {
		switch {
		case n == 1:
			return total, fmt.Errorf("err\n")
		case total == 0:
			return aux(prev-1, prev*(n-1))
		default:
			return aux(prev-1, prev*total)
		}
	}

	return aux(n, 0)
}

func main() {
	val, _ := fact(5)
	fmt.Println(val)
}
