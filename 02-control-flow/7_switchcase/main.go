package main

import "time"

func main() {
	age := 20

	switch age {
	case 10:
		println("You are 10 years old")
	case 11:
		println("You are 11 years old")
	case 17:
		println("You are 17 years old")
	case 18:
		println("You are eligible to vote")
	default:
		println("Your age is 20")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("This is enjoy day")
	case time.Monday:
		println("This is start of the week")
	case time.Thursday:
		println("This is almost weekend")
	default:
		println("This is a hard workiing day")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			println("I am an integer", t)
		case string:
			println("I am a string", t)
		case bool:
			println("I am a boolean", t)
		default:
			println("I am something else", t)
		}
	}
	whoAmI(21)
	whoAmI("Hello")
	whoAmI(true)
	whoAmI(3.14)
}
