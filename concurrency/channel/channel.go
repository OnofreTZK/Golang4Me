package main

import "fmt"

func main() {
	ch := make(chan int, 1) // second param is the buffer size

	ch <- 1 // sending to the channel of integers a integer.
	<-ch    // reading channel

	ch <- 2
	fmt.Println(<-ch)
}
