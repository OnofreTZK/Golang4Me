package main

import "fmt"

func main() {

	// maps must be initialized
	athletes := make(map[int]string) // key -> integer, value -> string

	athletes[0] = "Ronnie Coleman"
	athletes[1] = "Jay Cutler"
	athletes[3] = "Alex dos Santos"

	fmt.Println(athletes)

	for pos, name := range athletes {
		fmt.Printf("%s Position: %d\n", name, pos)
	}

	fmt.Println(athletes[0])
	delete(athletes, 1)
	fmt.Println(athletes[1])
}
