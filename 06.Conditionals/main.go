package main

import "fmt"

func main() {
	value1 := 1
	value2 := 2
	//OPERATORS >, <, >=, <=,==, && (AND), || (OR), ! (NOT)

	//SIMPLE CONDITIONS
	fmt.Println("SIMPLE CONTIONAL")
	if value1 == 1 {
		fmt.Println("VALUE IS 1")
	} else {
		fmt.Println("VALUE IS NOT 1")
	}
	//CONDITIONAL WITH AND (&&)
	fmt.Println("CONDITIONAL AND (&&)")
	if value1 == 1 && value2 == 2 {
		fmt.Println("VALUE1 ARE 1 AND VALUE2 ARE 2")
	}
	//CONDTIONAL OR (||)
	fmt.Println("CONDITIONAL OR (||)")
	if value1 == 1 || value2 >= 1 {
		fmt.Println("VALUE1 OR VALUE2 ARE 1 OR GREATER THAN 1")
	}
	//CONDTIONAL NOT (!)
	fmt.Println("CONDITIONAL NOT (!)")
	if !(value1 == 2) {
		fmt.Println("VALUE1 IS NOT 2")
	}
	//SWITCHS
	switch value1 {
	case 0:
		fmt.Println("VALUE1 IS 0")
	case 1:
		fmt.Println("VALUE1 IS 1")
	default:
		fmt.Println("VALUE1 IS NOT 0 OR 1")
	}
	//
	switch value3 := value1 * 3; value3 {
	case 0:
		fmt.Println("VALUE3 IS 0")
	case 3:
		fmt.Println("VALUE3 IS 3")
	default:
		fmt.Println("VALUE3 IS NOT 0 OR 3")
	}
	//SWITCHES WITH CONDITIONALS
	switch {
	case value1 == 0:
		fmt.Println("VALUE1 IS 0")
	case value1 == 1:
		fmt.Println("VALUE1 IS 1")
	default:
		fmt.Println("VALUE1 IS NOT 0 OR 1")
	}

}
