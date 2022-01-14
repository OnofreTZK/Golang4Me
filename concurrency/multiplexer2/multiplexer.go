package main

import (
	"fmt"
	"time"
)

func talk(person string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s talking: %d", person, i)
		}
	}()

	return c
}

func join(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-entry1:
				c <- s
			case s := <-entry2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := join(talk("Miriko"), talk("Hawk"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
