package main

import "fmt"

func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 2, 4, 5, 1, 6, 3}
	fmt.Println("Original slice:", numbers)

	uniqueNumbers := removeDuplicates(numbers)
	fmt.Println("Slice after removing duplicates:", uniqueNumbers)
}
