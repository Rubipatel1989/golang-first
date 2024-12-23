package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num + 30
}
func main() {
	// var num int
	// var ptr *int
	// ptr = &num

	num := 200
	ptr := &num

	// fmt.Println("Num has value:", num)
	fmt.Println("Pointers containes:", ptr)
	fmt.Println("Data contains through pointer:", *ptr)

	var pointer *int // Default pointer nill
	if pointer == nil {
		fmt.Println("There is not assigned")
	} else {
		fmt.Println("Has values")
	}
	value := 50
	modifyValueByReference(&value)
	fmt.Println("Value contain is:", value)

}
