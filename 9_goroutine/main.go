package main

import (
	"fmt"
	"time"
)

func printMesage(message string) {
	fmt.Println(message)
}

func main() {
	go printMesage("Hello from Goroutine!")
	fmt.Println("Hello from Main Function!")
	time.Sleep(2 * time.Second)
}
