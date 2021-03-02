package main

import "fmt"

func func1() {
	fmt.Println("Hello world!")
}
func func2(a, b int, c string) {
	fmt.Println(a, b, c)
}
func func3(a int) int {
	return a * 2
}

func func4(a int) (int, int) {
	return a, a * 2
}
func main() {
	//DEFAULT FUNCTIONS
	func1()
	//FUNCTION WITH PARAMETERS
	func2(1, 2, "Hello world!")
	//FUNCTION WITH RETURN ONE VALUE
	value := func3(2)
	fmt.Println("Value", value)
	//FUNCTION WITH RETURN MORE VALUES
	value1, _ := func4(2)
	fmt.Println("Value 1->", value1)
}
