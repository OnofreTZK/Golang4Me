package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":3000", nil))
	fmt.Println("vim-go")
}
