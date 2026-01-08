package main

import "fmt"

func ReservseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "ATAL"
	reverse := ReservseString(input)
	fmt.Println("Main string", input)
	fmt.Println("Reversed string", reverse)
}
