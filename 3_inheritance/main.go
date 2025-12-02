package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}
type Salary struct {
	Employee
	BasicPay float64
}

func (e Employee) GetDetails() {
	fmt.Println(e.Age, " ", e.EmployeeID, " ", e.Name)
}

func main() {
	salary := Salary{
		Employee: Employee{
			Person: Person{
				Name: "John Doe",
				Age:  30,
			},
			EmployeeID: 12345,
		},
		BasicPay: 50000.0,
	}

	fmt.Println("Name:", salary.Name)
	fmt.Println("Age:", salary.Age)
	fmt.Println("Employee ID:", salary.EmployeeID)
	fmt.Println("Basic Pay:", salary.BasicPay)

}
