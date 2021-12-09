package main

import "fmt"

func func1() {
	fmt.Println("Function1")
} // No parameter, no return

func func2(parameter1, parameter2 string) {
	fmt.Printf("Function2: p1: %s | p2: %s\n", parameter1, parameter2)
} // two parameters, no return

func func3() string {
	return "Function3"
} // No parameter, return a string

func func4(parameter1, parameter2 string) string {
	return fmt.Sprintf("Function4: p1: %s | p2: %s\n", parameter1, parameter2)
} // two parameters, return a string

func func5() (string, string) {
	return "String1", "String2"
}

func main() {
	func1()
	func2("Param1", "Param2")
	fmt.Println(func3())
	fmt.Println(func4("Param1", "Param2"))
	str1, str2 := func5()
	fmt.Println(str1, str2)

}
