package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Integers
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal Integer is", reflect.TypeOf(32000))

	// Unsigned Integers -> uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("The byte is", reflect.TypeOf(b))

	// Signed Integers -> int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("The max value for a 64 bits integer is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	// Rune is a mapping of the Unicode table (int32)
	var i2 rune = 'a'
	fmt.Println("Rune is", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Float numbers -> float32, float64
	var x float32 = 49.99
	fmt.Println("The type of x is", reflect.TypeOf(x))
	fmt.Println("Literal type of 49.99 is", reflect.TypeOf(49.99)) //float64

	// Booleans
	bo := true
	fmt.Println("The type of bo is", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// Strings
	str1 := "This is a string"
	fmt.Println(str1 + " Concatened")
	fmt.Println("The size of the string is", len(str1))

	str2 := `This
	is 
	as
	string`
	fmt.Println("The size of the string with tabs is", len(str2))

	// Char -> Golang don't have char type. it is a int32
	char := 'a'
	fmt.Println("The type of char is", reflect.TypeOf(char))
	fmt.Println(char)
}
