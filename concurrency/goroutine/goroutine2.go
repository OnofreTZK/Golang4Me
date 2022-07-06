package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello, World") // goroutine
	// it wont wait to end to execute the next instruction
	write("Coding in Go")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
