package main

import (
	"fmt"
	"time"
)

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Unknown type"
	}
}

func main() {
	fmt.Println(getType(2.3))
	fmt.Println(getType(1))
	fmt.Println(getType("STR"))
	fmt.Println(getType(func() {}))
	fmt.Println(getType(time.Now()))
}
