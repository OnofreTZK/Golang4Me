package main

import (
	"cod3rcourse/htmltitle"
	"fmt"
	"time"
)

func fastest(url1, url2, url3 string) string {
	c1 := htmltitle.Title(url1)
	c2 := htmltitle.Title(url2)
	c3 := htmltitle.Title(url3)

	// control structure for concurrency -> pretty much like a switch case
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "They're all to slow!"
		//default:
		//	return "No answers yet"
	}
}

func main() {

	champs := fastest(
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
	)

	fmt.Println(champs)
}
