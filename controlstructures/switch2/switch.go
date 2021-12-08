package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { // boolean switch -> the first true case
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}
}
