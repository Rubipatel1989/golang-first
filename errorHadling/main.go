package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("demoninator cannot be zero")
	}
	return a / b, nil
}
func main() {
	// fmt.Println("Started error handling")
	// Divide, error := divide(10, 0)
	// if error != nil {
	// 	fmt.Println(error)
	// }
	// fmt.Println(Divide)

	fmt.Println("Started error handling")
	Divide, _ := divide(10, 0)
	fmt.Println(Divide)

}
