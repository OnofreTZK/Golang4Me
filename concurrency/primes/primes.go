package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primes(n int, c chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				c <- prime
				start = prime + 1
				time.Sleep(time.Millisecond * 200)
				break // because the for loop here is working like a while(true)
			}
		}
	}

	close(c) // avoid deadlock
}

func main() {
	ch := make(chan int, 30)
	go primes(cap(ch), ch)
	// all of this just to show we can use for range in a channel
	for i := range ch {
		fmt.Printf("%d ", i)
	}
	fmt.Println("End")
}
