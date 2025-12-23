package main

import (
	"fmt" // Import the fmt package for correct printing
)

// reverseArray reverses a slice of integers in place.
func reverseArray(arr []int) []int {
	start := 0
	end := len(arr) - 1
	for start < end {
		// Swap elements using Go's concise syntax
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}

func main() {
	num := []int{10, 20, 30, 40, 50}

	// The slice is reversed in place by the function
	reversed := reverseArray(num)

	// 💡 Use fmt.Println to print the slice content correctly
	fmt.Println("Reversed array is ", reversed)
}
