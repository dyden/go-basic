package main

import (
	"fmt"
	"time"
)

func say(message string, c chan<- string) {
	time.Sleep(time.Second * 5)
	c <- message
}

func main() {
	//channel<- INPUT    <-channel  OUTPUT
	c := make(chan string, 1)
	go say("Hello", c)
	fmt.Println("world!")
	fmt.Println(<-c)
}
