package main

import "fmt"

//PC STRUCT
type pc struct {
	brand string
	model string
	price int
}

func (myPC *pc) String() string {
	return fmt.Sprintf("Brand: %s\nModel: %s\nPrice: %d", myPC.brand, myPC.model, myPC.price)
}

func main() {
	//CREATE NEW PC
	myPC := pc{brand: "Asus", model: "ROG", price: 1000}
	//PRINT PC
	fmt.Println(myPC)
	//PRINT WITH CUSTOM STRING
	fmt.Println("NEW FORMAT TO PRINT PC")
	fmt.Println(myPC.String())

}
