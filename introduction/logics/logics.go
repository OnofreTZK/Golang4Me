package main

import "fmt"

func buy(job1, job2 bool) (bool, bool, bool) {

	fazer250abs := job1 && job2
	ipadPro := job1 != job2 // gambiarra for exclusive or
	burger := job1 || job2

	return fazer250abs, ipadPro, burger
}

func main() {
	fazer, ipad, burger := buy(true, true)
	fmt.Printf("Fazer: %t, Ipad: %t, Burger: %t, Healthy: %t\n", fazer, ipad, burger, !burger)
}
