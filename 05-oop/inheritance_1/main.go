package main

import "strconv"

type Education struct {
	Name    string
	Age     int
	Degree  string
	College string
	Streem  string
	Year    int
}

func (e Education) GetEducationDetails() string {
	return "Dear " + e.Name + " with age of " + strconv.Itoa(e.Age) + " Years, You passed " + e.Degree + " in " + e.Streem + " from " + e.College + ", Year: " + strconv.Itoa(e.Year)
}

func main() {
	edu := Education{
		Name:    "Pawan Kumar",
		Age:     36,
		Degree:  "Bachelor of Science",
		College: "UPTU",
		Streem:  "Information Technology",
		Year:    2010,
	}

	details := edu.GetEducationDetails()
	println(details)
}
