package main

import "fmt"

func gradeForConcept(n float64) string {

	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 7 && n < 9 {
		return "B"
	} else if n >= 5 && n < 7 {
		return "C"
	} else {
		return "D"
	}

}

func main() {
	fmt.Println(gradeForConcept(9.5))
	fmt.Println(gradeForConcept(8))
	fmt.Println(gradeForConcept(5.8))
	fmt.Println(gradeForConcept(3.1))
}
