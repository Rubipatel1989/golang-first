package main

import "strconv"

type Education struct {
	Degree  string
	College string
	Streem  string
	Year    int
}

func (e Education) GetEducationDetails() string {
	return e.Degree + " in " + e.Streem + " from " + e.College + ", Year: " + strconv.Itoa(e.Year)
}

func main() {
	edu := Education{
		Degree:  "Bachelor of Science",
		College: "XYZ University",
		Streem:  "Computer Science",
		Year:    2022,
	}

	details := edu.GetEducationDetails()
	println(details)
}
