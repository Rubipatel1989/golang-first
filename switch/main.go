package main

import "fmt"

func main() {
	day := 61
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	month := 3
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	case 2:
		fmt.Println("28 days")
	default:
		fmt.Println("Invalid month")
	}

	temperature := 30
	switch {
	case temperature < 0:
		fmt.Println("It's freezing")
	case temperature >= 0 && temperature <= 10:
		fmt.Println("It's cold")
	case temperature > 10 && temperature <= 20:
		fmt.Println("It's cool")
	case temperature > 20 && temperature <= 30:
		fmt.Println("It's warm")
	default:
		fmt.Println("It's hot")
	}

	year := 2000
	switch {
	case year > 1000:
		fmt.Println("This is the 2nd millennium")
	
	}

}
