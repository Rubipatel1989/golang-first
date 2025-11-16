package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d end\n", i)
}

//worker 1 started

func main() {
	//fmt.Println("Exploure goroutne started")
	var wg sync.WaitGroup
	// Start 3 workers concurrently
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}
	wg.Wait() // Wait for all workers to finish
	fmt.Println("workers task completed")
}
