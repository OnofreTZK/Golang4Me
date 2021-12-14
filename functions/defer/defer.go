package main

import "fmt"

func getValue(num int) int {
	defer fmt.Println("Game over!")

	if num >= 5000 {
		fmt.Println("High Value!")
		return 5000
	} else {
		fmt.Println("Low Value")
		return num
	}
}

func main() {
	fmt.Println(getValue(6000))
	fmt.Println(getValue(4000))
}
