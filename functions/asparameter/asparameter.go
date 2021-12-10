package main

import "fmt"

func multi(a, b int) int {
	return a * b
}

func exec(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func main() {
	fmt.Println(exec(multi, 2, 6))
}
