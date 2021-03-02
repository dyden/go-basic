package main

import (
	"fmt"

	carpackage "github.com/dyden/go-basic/11.Data_access_modifiers/models" // YOU HAVE TO USE 'go mod init '...' IN THE ROOT FOLDER
)

func main() {
	var myCard carpackage.PublicCar
	myCard.Brand = "Fiat"
	myCard.Year = 2018
	fmt.Println(myCard)
	// var myOtherCard carpackage.privat.... //YOU CANNOT ACCESS TO PIVATE STRUCT

	//PUBLIC FUNCTION
	carpackage.PublicFunction()

}
