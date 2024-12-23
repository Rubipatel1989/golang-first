package main

import "fmt"

func main() {

	x := 101
	if x == 100 {
		fmt.Println("x is 100")
	} else if x > 100 {
		fmt.Println("x is greater than 100")
	} else {

		fmt.Println("x is smaller than 100")
	}
	y := 200
	z := 500
	if x == 100 && (y == 200 || z < 43) {
		fmt.Println("x is 100 and y is 200")
	} else {
		fmt.Println("x is not 100 and y is not 200")
	}
}
