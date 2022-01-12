package main

import "fmt"

func routine(c chan int) {
	fmt.Println("Executed!")
	c <- 1
	c <- 2
	c <- 3
	fmt.Println("Buffer is full here!")
	c <- 4
	fmt.Println("Will be executed after one element in the buffer is consumed")
	c <- 5
	fmt.Println("This string will never be printed")
	c <- 6
	c <- 7
}

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	fmt.Println(<-ch)
}
