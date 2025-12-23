package main

import (
	"fmt"
	"sync"
)

// A Goroutine is a lightweight, independent function execution
// that runs concurrently with other functions in a Go program.

// 1. They are Go's way of achieving concurrency.
// 2. They are much cheaper to create than traditional operating system threads.
// 3. Goroutines are managed by the Go runtime and multiplexed onto a small number of OS threads.
// 4. They communicate primarily using channels.

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}
func sayHi() {
	fmt.Println("Hi")
}

func main() {
	// Start two goroutines
	//go sayHello()
	//sayHi()
	// wait := make(chan struct{})
	// go func() {
	// 	sayHello()
	// 	close(wait)
	// }()
	// <-wait
	// wait := make(chan struct{})
	// go func() {
	// 	sayHello()
	// 	wait <- struct{}{}
	// }()
	// <-wait
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)
	sayHi()
	wg.Wait()
}
