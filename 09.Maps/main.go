package main

import "fmt"

func main() {
	//MAP IS A COLLECTION OF KEY-VALUE PAIRS
	persons := make(map[string]int) //NAME-AGE
	persons["Jhon"] = 18
	persons["Mary"] = 20
	persons["Mike"] = 15
	fmt.Println(persons)
	//PRINT EACH KEY-VALUE PAIR
	for key, value := range persons {
		fmt.Printf("Name: %s, Age: %d\n", key, value)
	}

	//FIND SOME KEY
	fmt.Println("FINDING 'Jhon'")
	value := persons["Jhon"]
	fmt.Println(value)
	fmt.Println("FINDING 'Jhonny'")
	value2 := persons["Jhonny"]
	fmt.Println(value2) //IF NOT FOUND, RETURN 0

	//DELETE A KEY-VALUE PAIR
	fmt.Println("REMOVING 'Jhon' OF THE MAP")
	delete(persons, "Jhon")
	fmt.Println(persons)

	//ADD A NEW KEY-VALUE PAIR
	fmt.Println("ADDING 'Petter' OF THE MAP")
	persons["Petter"] = 25
	fmt.Println(persons)
}
