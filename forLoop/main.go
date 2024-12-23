package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Number is ", i)
	}
	counter := 0
	for { // Infinite loop
		fmt.Println("Counter is ", counter)
		counter++
		if counter == 20 {
			break
		}
	}

	numbers := []int{1, 22, 3, 4, 5, 6, 7, 88, 91, 10}
	for index, value := range numbers {
		fmt.Printf("Index of number is %d, and value is %d\n", index, value)
	}

	data := "Love You"
	for index, char := range data {
		fmt.Printf("index of data is %d, and char is %c\n", index, char)
	}
}
