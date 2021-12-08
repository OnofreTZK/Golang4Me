package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // while
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { // normal for
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMixing...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Even ")
		} else {
			fmt.Print("Odd ")
		}
	}

	for { // infinite loop
		fmt.Println("For ever...")
		time.Sleep(time.Second)
	}

}
