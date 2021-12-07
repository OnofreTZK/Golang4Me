package main

import "fmt"

func main() {

	i := 1

	// Go doesn't accept pointer arithmetic

	var p *int = nil

	p = &i  // i address in memory
	h := *p // value pointed by p

	fmt.Println(p, *p, i, h, &i)
}
