package main

import "fmt"

// Each go file can afford a init func
func init() {
	fmt.Println("Initializing")
}

func main() {
	fmt.Println("Main")
}
