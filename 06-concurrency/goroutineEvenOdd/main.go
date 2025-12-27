package main

import (
	"fmt"
	"sync"
)

var evens []int
var odds []int

func findEvens(nums []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
}

func findOdds(nums []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range nums {
		if num%2 != 0 {
			odds = append(odds, num)
		}
	}
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	wg.Add(2)
	go findEvens(nums, &wg)
	go findOdds(nums, &wg)

	wg.Wait()

	fmt.Println("Even numbers", evens)
	fmt.Println("Odds numbers", odds)
}
