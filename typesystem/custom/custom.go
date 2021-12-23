package main

import "fmt"

type grade float64

func (g grade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func gradeForConcept(g grade) string {

	if g.between(9.0, 10.0) {
		return "A"
	} else if g.between(7.0, 8.9) {
		return "C"
	} else if g.between(5.0, 6.9) {
		return "C"
	} else {
		return "D"
	}

}

func main() {

	fmt.Println(gradeForConcept(7.6))
}
