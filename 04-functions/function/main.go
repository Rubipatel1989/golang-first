package main

import (
	"fmt"
)

func simpleFunction() {
	fmt.Println("Ajay Kumar Good")
}
func add(a int, b int) int {
	return a + b
}
func main() {
	fmt.Println("Pawan Kumar Good")
	simpleFunction()
	Answer := add(12, 13)
	fmt.Println(Answer)
}
