package main

import "fmt"

func getResult(grade float64) string {

	// ok google...
	if grade >= 6 {
		return "Apavorou"
	}
	return "Macetado"
	// return grade >= 6 ? "Apavorou" : "Macetado"

}

func main() {
	fmt.Println(getResult(6.2))
	fmt.Println(getResult(5.9))
}
