package main

import "fmt"

//PC STRUCT
type pc struct {
	brand string
	model string
	price int
}

func (nyPC *pc) dublicatePrice() {
	nyPC.price = nyPC.price * 2
}
func main() {
	//POINTERS
	a := 10
	b := &a
	fmt.Println(a)
	fmt.Println(b)  //IF WE PRINT 'b' WE GET THE MEMORY ADDRESS
	fmt.Println(*b) //IF WE PRINT '*b' WE GET THE VALUE OF THE MEMORY ADDRESS
	//CHAGE VALUE ON MEMORY ADDRESS
	*b = 20 //NOW WE CHANGE THE VALUE OF 'a' TO 20
	fmt.Println(a)

	//POINTERS ON STRUCTS
	myPC := pc{brand: "DELL", model: "XPS", price: 1000}
	fmt.Println(myPC)
	//DUPLICATE THE PRICE OF THE PC WITH THE POINTER
	myPC.dublicatePrice()
	fmt.Println(myPC)

}
