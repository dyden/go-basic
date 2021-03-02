package main

import "fmt"

//NEW STRUCT
type car struct {
	brand string
	year  int
}

func main() {
	//ONE WAY TO INSTANTIATE A STRUCT
	myCard := car{brand: "Ford", year: 2020}
	fmt.Println(myCard)
	//ANOTHER WAY TO INSTANTIATE A STRUCT
	var otherCard car
	otherCard.brand = "Toyota"
	otherCard.year = 2019
	fmt.Println(otherCard)
}
