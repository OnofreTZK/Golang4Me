package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNum() int {
	seed := rand.NewSource(time.Now().UnixNano()) // Current nano second
	r := rand.New(seed)

	return r.Intn(10)
}

func main() {

	if i := randomNum(); i > 5 {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}
}
