package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is", num)
	fmt.Println("Type of number is", fmt.Sprintf("%T", num))
	var data float64 = float64(num)
	data = data + 2.25
	fmt.Println("Data is", data)
	fmt.Println("Type of data is", fmt.Sprintf("%T", data))

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Data is", str)
	fmt.Println("Type of data is", fmt.Sprintf("%T", str))

	numberString := "123"
	number_int, _ := strconv.Atoi(numberString)
	number_int = number_int + 2
	fmt.Println("number_int is", number_int)
	fmt.Println("Type of number_int is", fmt.Sprintf("%T", number_int))

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)

	fmt.Println("number_float is", number_float)
	fmt.Println("Type of number_float is", fmt.Sprintf("%T", number_float))
}
