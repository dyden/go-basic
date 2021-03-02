package main

import (
	"fmt"
	"sync"
	"time"
)

func say(message string, wg *sync.WaitGroup) {
	defer wg.Done() // TASK DONE AT THE END OF THE FUNCTION
	//AWAIT FOR 10 SECONDS
	time.Sleep(time.Second * 5)
	fmt.Println("Hello, world!")
}

func main() {
	var wg sync.WaitGroup // CREATE A WAIT GROUP
	wg.Add(1)             // ADD 1 TASK TO THE WAIT GROUP
	go say("Hello", &wg)  // RUN THE FUNCTION AS A GOROUTINE
	fmt.Println("world!")
	wg.Wait() // WAIT FOR THE TASK TO FINISH
}
