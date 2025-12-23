package main

// A Go channel is a built-in data structure that allows two goroutines
// to communicate and synchronize their execution.

// 1. It acts as a conduit for a specific data type.
// 2. Data is passed between goroutines using send and receive operations.
// 3. Channels can be unbuffered (blocking until the receiver is ready)
// 	or buffered (allowing a limited number of values to be stored).
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
