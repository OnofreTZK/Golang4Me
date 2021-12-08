package main

import "fmt"

func gradeForConcept(n float64) string {

	var grade = int(n)

	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
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
