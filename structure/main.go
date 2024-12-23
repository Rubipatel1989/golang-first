package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}
type Address struct {
	House int
	Area  string
	State string
}
type Employee struct {
	Persion_Details person
	Persion_Contact Contact
	Persion_Address Address
}

func main() {
	// 1st Method
	var prince person
	prince.FirstName = "Prince"
	prince.LastName = "Kumar"
	prince.Age = 45

	fmt.Println(prince)
	// 2nd Method
	p1 := person{
		FirstName: "Pawan",
		LastName:  "Kumar",
		Age:       22,
	}
	p2 := person{
		FirstName: "Saroj",
		LastName:  "Kumar",
		Age:       25,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	var employee1 Employee
	employee1.Persion_Details = person{
		FirstName: "Pawan",
		LastName:  "Kumar",
		Age:       21,
	}
	employee1.Persion_Contact.Email = "pawankum0010@gmail.com"
	employee1.Persion_Contact.Phone = "8010043297"
	employee1.Persion_Address = Address{
		House: 213,
		Area:  "Very Long Are",
		State: "Jharkand",
	}
	fmt.Println(employee1)
}
