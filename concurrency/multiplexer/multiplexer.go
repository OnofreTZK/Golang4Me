package main

import (
	"cod3rcourse/htmltitle"
	"fmt"
)

func forward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

// multiplex - join the messages in a channel
func join(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(entry1, c)
	go forward(entry2, c)
	return c
}

func main() {
	c := join(
		htmltitle.Title("https://www.cod3r.com.br", "https://www.google.com"),
		htmltitle.Title("https://www.amazon.com.br", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
