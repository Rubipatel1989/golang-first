package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time", currentTime)
	fmt.Printf("Type of time %T\n", currentTime)

	formatedDate := currentTime.Format("02-01-2006, 03:04 PM")
	fmt.Println("Changed format is", formatedDate)

	dateStr := "2024-12-23"
	layoutStr := "2006-01-02"
	formatedStr, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("Formated Date is", formatedStr)

}
