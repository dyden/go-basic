package carpackage

import "fmt"

//IF U USE LOWERCASE, IT IS PRIVATE
type privatelicCar struct {
	brand string
	year  int
}

// IF U USE UPPERCASE, IT IS PUBLIC
//PublicCar is a public struct
type PublicCar struct {
	Brand string
	Year  int
	// year int			IF USE USE LOWERCASE, THIS FIELD WILL BE PRIVATE  AS WELL
}

//PublicFunction is a public function
func PublicFunction() {
	fmt.Println("Hello world!")
}
