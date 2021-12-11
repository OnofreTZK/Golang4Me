package main

import "fmt"

func printApproveds(approveds ...string) {
	for i, appr := range approveds {
		fmt.Printf("%d) %s\n", i+1, appr)
	}
}

func main() {
	approveds := []string{"Shayera Hol", "Adelaine Bordeaux", "Ka'Toye Orari"}
	printApproveds(approveds...) // Spread the values in the slice as arguments in the function
	// Only works with slice
}
