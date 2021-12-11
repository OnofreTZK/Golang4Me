package main

import "fmt"

func mean(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Mean: %.2f\n", mean(7.7, 8.9, 5.6))
	fmt.Printf("Mean: %.2f\n", mean(7.7, 8.9))
}
