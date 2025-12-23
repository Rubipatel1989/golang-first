package main

import "fmt"

func main() {

	age := 22
	name := "Pawan Kumar"
	height := 5.8
	fmt.Println("Age:", age, "Name:", name, "Height:", height)
	fmt.Print("Age:", age, "Name:", name, "Height:", height)
	fmt.Printf("Age:%d\n Name:%s\n Height:%0.2f\n", age, name, height)
	fmt.Printf("Type of age:%T\n", age)
}
