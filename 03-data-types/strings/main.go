package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Apple,Mango,Banana,Orange,Pineapple"
	parts := strings.Split(data, ",")
	fmt.Println(parts[0])

	str := "I am the best for 2 days for pawan for sonu"
	count := strings.Count(str, "for")
	fmt.Println("Count is", count)

	str = "   Hello,    Go"
	trimmed := strings.TrimSpace(str)
	fmt.Println("Triimed values are", trimmed)

	str1 := "Mani"
	str2 := "Kumar"
	result := strings.Join([]string{str1, str2}, ", ")
	fmt.Println(result)

}
