package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	// Identifier with Capital letter is public
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := product{1, "iPad Pro 11", 7000.00, []string{"Tablet", "Apple"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json to struct
	var p2 product
	jsonString := string(p1Json)
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
