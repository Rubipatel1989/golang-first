package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p People) Intro() {
	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}

func main() {
	person := People{
		Name: "Alice",
		Age:  28,
	}

	person.Intro()
}
