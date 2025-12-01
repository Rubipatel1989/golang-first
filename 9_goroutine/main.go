package main

// A Goroutine is a lightweight, independent function execution
// that runs concurrently with other functions in a Go program.

// 1. They are Go's way of achieving concurrency.
// 2. They are much cheaper to create than traditional operating system threads.
// 3. Goroutines are managed by the Go runtime and multiplexed onto a small number of OS threads.
// 4. They communicate primarily using channels.

import (
	"fmt"
	"time"
)

func printMesage(message string) {
	fmt.Println(message)
}

func main() {
	go printMesage("Hello from Goroutine!")
	fmt.Println("Hello from Main Function!")
	time.Sleep(2 * time.Second)
}
