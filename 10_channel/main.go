package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numchan chan int) {
	for num := range numchan {
		fmt.Println("Processing number...", num)
		time.Sleep(300000 * time.Microsecond)
	}

}
func main() {
	numChan := make(chan int)
	go processNum(numChan)

	for {
		numChan <- rand.Intn(10)
	}

	// messageChannel := make(chan string)
	// messageChannel <- "Hello, Channels!" // This will cause a deadlock [Blocking operation]
	// msg := <-messageChannel
	// fmt.Println(msg)

}
