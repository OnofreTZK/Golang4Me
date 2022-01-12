package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // block ops
	fmt.Println("After block ops")
}

func main() {
	c := make(chan int) // Channel without buffer

	go routine(c)

	fmt.Println(<-c) // block ops
	fmt.Println("read")
	fmt.Println(<-c) // deadlock because there's no instruction to add data to this channel anymore
	// all go routines are dead here
	fmt.Println("End")
}
