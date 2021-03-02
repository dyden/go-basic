package main

import "fmt"

func main() {
	//DEFER ----> IS USED TO EXECUTE A FUNCTION AT THE END OF A FUNCTIONS
	defer fmt.Println("Hello") //THIS LINE WILL BE EXECUTED AT THE END OF THE MAIN FUNCTION
	fmt.Println("world!")

	//CONTINUE AND BREAK
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue //	WE USE CONTINUE TO SKIP THE CURRENT ITERATION
		}
		if i == 5 {
			break // WE USE 'BREAK' FOR STOP THE LOOP
		}
		fmt.Println(i)
	}
}
