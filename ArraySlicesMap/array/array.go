package main

import "fmt"

func main() {

	// arrays are homogenic and static
	var grades [3]float64
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.8, 4.9, 9.0
	fmt.Println(grades)

	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	mean := total / float64(len(grades))
	fmt.Printf("Mean: %.2f\n", mean)

}
