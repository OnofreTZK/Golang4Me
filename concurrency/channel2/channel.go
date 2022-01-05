package main

import (
	"fmt"
	"time"
)

// Channel -> Communication way between goroutines
// Channel is a type

func twoThreeFourTimes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go twoThreeFourTimes(2, c)
	fmt.Println("waiting A")

	// This line will wait the two async data to be ready and only then keep going.
	a, b := <-c, <-c // reading datas from channel
	fmt.Println("waiting B")
	fmt.Println(a, b)

	fmt.Println("waiting C")
	fmt.Println(<-c)
}
