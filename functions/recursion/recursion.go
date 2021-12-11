package main

import (
	"fmt"
)

func fact(n int) (int, error) {

	var aux func(int, int) (int, error)

	aux = func(prev, total int) (int, error) {
		fmt.Printf("Total: %d | n: %d\n", total, prev)
		switch {
		case prev == 1:
			return total, fmt.Errorf("err\n")
		case prev == n:
			return aux(prev-1, prev*total)
		default:
			return aux(prev-1, (prev-1)*total)
		}
	}

	return aux(n, n-1)
}

func main() {
	val, _ := fact(10)
	fmt.Println(val)
}
