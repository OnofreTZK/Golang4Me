package main

import "fmt"

func printResult(grade float64) {

	if grade >= 7 {
		fmt.Println("Macetou com nota", grade)
	} else {
		fmt.Println("Destruido com nota", grade)
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
}
