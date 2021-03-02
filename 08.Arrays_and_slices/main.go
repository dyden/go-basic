package main

import "fmt"

func main() {
	//THE MAIN DIFERENCE BETWEN 'ARRAY' AND 'SLICE' IS THAT THE SIZE OF 'SLICE' IS MUTABLE AND THE SIZE OF 'ARRAY' IS NOT MUTABLE
	//ARRAY
	var array [4]int //EMPTY ARRAY [0,0,0,0]
	fmt.Println(array)
	//INIT ARRAY
	array[0] = 0
	array[1] = 1
	array[2] = 2
	array[3] = 3

	fmt.Println(array)
	//LENGHT OF ARRAY
	fmt.Printf("LENGTH OF ARRAY: %d\n", len(array))
	//MAX CAPACITY OF ARRAY
	fmt.Printf("MAX CAPACITY OF ARRAY: %d\n", cap(array))
	//SLICE
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice)
	//LENGTH OF SLICE
	fmt.Printf("LENGTH OF SLICE: %d\n", len(slice))
	//MAX CAPACITY OF SLICE
	fmt.Printf("MAX CAPACITY OF SLICE: %d\n", cap(slice))

	//PRINT VALUES OF BOTH
	fmt.Println("FIRST VALUE OF SLICE")
	fmt.Println(slice[0])
	fmt.Println("FIRST 3 VALUES OF ARRAY")
	fmt.Println(array[:3])
	fmt.Println("FROM 2ND VALUE TO END OF SLICE")
	fmt.Println(slice[2:])
	fmt.Println("FROM 2ND VALUE OF ARRAY TO 4TH VALUE")
	fmt.Println(array[1:4])

	//ADD NEW VALUES TO SLICE
	fmt.Println("ADDING NEW DATA TO SLICE")
	slice = append(slice, 6)
	fmt.Println(slice)
	//ADD NEW SLICE TO CURENT SLICE
	fmt.Println("ADDING NEW SLICE TO CURRENT SLICE")
	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//REMOVE FIST VALUE FROM SLICE
	fmt.Println("REMOVING FIRST VALUE FROM SLICE")
	slice = append(slice[1:])
	fmt.Println(slice)
	fmt.Println("REMOVING FIRST 3 VALUES FROM SLICE")
	slice = append(slice[3:])
	fmt.Println(slice)
	fmt.Println("REMOVING RANGE OF VALUES FROM SLICE")
	slice = append(slice[1:5])
	fmt.Println(slice)

	// GETTING VALUES OF SLICE WITH LOOP
	for _, value := range slice {
		fmt.Println(value)
	}

}
