package main

import "fmt"

func main() {

	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println("FOR CONTITIONAL")
	//FOR CODITIONAL
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("FOR WHILE")
	//FOR WHILE
	counter := 0
	for counter <= 20 {
		fmt.Println(counter)
		counter += 2
	}
	fmt.Println("FOR RANGE")
	//FOR RANGE
	for _, value := range numbers {
		fmt.Println(value)
	}
	//INFINITE LOOP

	for {
		fmt.Println("INFINITE LOOP") // IF YOU DONT STOP IT, IT WILL RUN FOREVER
	}

}
