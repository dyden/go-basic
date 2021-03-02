package main

import "fmt"

func main() {
	//CONSTANTS
	fmt.Printf("CONSTANTS\n\n")

	const const1 float32 = 3.1415
	fmt.Printf("%f -> %T\n", const1, const1)
	//IF YOU DONT SPECIFIY ANY TYPE, IN THIS CASE GO WILL ASSUME IT IS A FLOAT 64
	const const2 = 3.1415
	fmt.Printf("%f -> %T\n", const2, const2)
	//IF YOU DONT SPECIFIY ANY TYPE, IN THIS CASE GO WILL ASSUME IT IS A STRING
	const const3 = "Hello"
	fmt.Printf("%s -> %T\n\n", const3, const3)

	//VARIABLES
	fmt.Printf("VARIABLES\n\n")
	var1 := 12 //----> IF YOU DONT SPECIFY ANY TYPE GO WILL ASSUME A DEFUALT TYPE (INT,STRING,FLOAT64...)
	var var2 int = 12
	var var3 bool = true
	var var4 string = "world"
	fmt.Printf("var1-> %d, var2->%d, var3->%t, var4->%s\n", var1, var2, var3, var4)
}
