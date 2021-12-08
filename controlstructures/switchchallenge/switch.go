package main

import "fmt"

func gradeForConcept(n float64) string {

	var grade = int(n)

	switch {
	case grade >= 9 && grade <= 10:
		return "A"
	case grade >= 7 && grade < 9:
		return "B"
	case grade >= 5 && grade < 7:
		return "C"
	case grade >= 3 && grade < 5:
		return "D"
	case grade >= 0 && grade < 3:
		return "E"
	default:
		return "Invalid grade!"
	}

}

func main() {
	fmt.Println(gradeForConcept(9.5))
	fmt.Println(gradeForConcept(8))
	fmt.Println(gradeForConcept(5.8))
	fmt.Println(gradeForConcept(3.1))
	fmt.Println(gradeForConcept(11))
}
