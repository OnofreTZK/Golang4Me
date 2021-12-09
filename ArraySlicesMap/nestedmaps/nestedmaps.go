package main

import "fmt"

func main() {

	namesByLetter := map[string]map[string]float64{ // key -> string, value -> map(string, float64)
		"G": {
			"Gabi":  10000.0,
			"Gotze": 100.0,
		},
		"V": {
			"Van Tyson": 1000000000000000.0,
			"Viktor":    1000.0,
		},
		"L": {
			"Ludwig":   199992820.0,
			"Lorraine": 1050.0,
		},
	}

	delete(namesByLetter, "L")

	for letter, names := range namesByLetter {
		fmt.Println(letter, names)
	}
}
