package main

import "fmt"

func main() {

	namesAndSalarys := map[string]float64{
		"Lebron James":  12222222.0,
		"Stephen Curry": 122233444.0,
		"Kevin Durant":  3999999999.6,
	}

	namesAndSalarys["Jamorango"] = 90000000000.2
	delete(namesAndSalarys, "Unknown")

	for name, salary := range namesAndSalarys {
		fmt.Println(name, salary)
	}
}
