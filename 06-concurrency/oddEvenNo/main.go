package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var evens []int
	var odds []int

	var wg sync.WaitGroup
	wg.Add(2)
	// Even number goroutine
	go func() {
		defer wg.Done()
		for _, num := range nums {
			if num%2 == 0 {
				evens = append(evens, num)
			}
		}
	}()
	// Odd number goroutine
	go func() {
		defer wg.Done()
		for _, num := range nums {
			if num%2 != 0 {
				odds = append(odds, num)
			}
		}
	}()
	wg.Wait()
	fmt.Println("Even Numbers:", evens)
	fmt.Println("Odd Numbers:", odds)
}
