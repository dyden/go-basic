package main

import "fmt"

func main() {
	//WE USE THE PACKAGE "fmt" TO PRINT IN THE CONSOLE
	fmt.Println("Hello World")
	//WE CAN USE 'fmt.printf' TO PRINT IN THE CONSOLE WITH A FORMAT
	age := 24
	fmt.Printf("My age is %d\n", age)
	/*
		BOOL:                    	%t
		INT, INT8 ETC..:         	%d
		UINT, UINT8 ETC..:	        %d, %#x IF PRINTER WITH %#v
		FLOAT32, COMPLEX64, ETC..:	%g
		STRING:                  	%s
		CHAN:	                    %p
		POINTER:	                %p
		....
		MORE ABOUT 'fmt': https://pkg.go.dev/fmt
	*/
}
