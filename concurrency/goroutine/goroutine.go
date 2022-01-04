package main

import (
	"fmt"
	"time"
)

func talk(person, text string, amount int) {

	for i := 0; i < amount; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {

	// Serial
	fmt.Println("Serial code")
	talk("Jolyne Cujoh", "Yare Yare Dawa...", 3)
	talk("Jotaro Cujoh", "Yare Yare Daze...", 1)

	// Threads or goroutines
	fmt.Println("With threads/goroutine")
	go talk("Jolyne Cujoh Parallel", "Yare Yare Dawa...", 10)
	talk("Jotaro Cujoh", "Yare Yare Daze...", 5)

}
