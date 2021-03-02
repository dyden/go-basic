package main

import "fmt"

func message(message string, c chan<- string) {
	c <- message
}

func main() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	fmt.Printf("LENGTH:%d CAP:%d\n", len(c), cap(c))
	//CLOSE -> WE USE CLOSE TO CLOSE THE CHANNEL
	close(c)
	//RANGE -> WE USE RANGE TO READ THE VALUES FROM THE CHANNEL
	for message := range c {
		fmt.Println(message)
	}
	//SELECT
	c2 := make(chan string)
	c3 := make(chan string)
	go message("messaje2", c2)
	go message("messaje3", c3)
	for i := 0; i < 2; i++ {
		select {
		case m2 := <-c2:
			fmt.Println("MESSAEJE RECEIVED FROM:", m2)
		case m3 := <-c3:
			fmt.Println("MESSAEJE RECEIVED FROM:", m3)

		}

	}

}
