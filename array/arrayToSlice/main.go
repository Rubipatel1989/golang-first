package main

import "fmt"

func main() {
	array1 := [5]int{10, 20, 30, 40, 50}
	slice := array1[:] // Slicing from index 1 to 3
	fmt.Println(slice)

}
