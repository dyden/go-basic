package main

import "fmt"

//INTERFACE
type shape interface {
	area() float64
}

//STRUCTS
type square struct {
	height float64
	width  float64
}
type rectangle struct {
	height float64
	width  float64
}

//FUNCTIONS
func (s square) area() float64 {
	return s.height * s.width
}
func (r rectangle) area() float64 {
	return r.height * r.width
}
func calculateArea(shape shape) {
	fmt.Println("Area:", shape.area())
}

func main() {
	//SQUARE
	s := square{height: 3, width: 3}
	r := rectangle{height: 2, width: 5}
	calculateArea(s)
	calculateArea(r)

	//LIST OF INSTERFACES
	myInterface := []interface{}{"Hi", 12, true, 1.2}
	fmt.Println(myInterface)
}
