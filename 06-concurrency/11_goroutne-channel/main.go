package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2 * time.Second)
	fmt.Println("Sayhello function after 5 seconds")
}
func sayHi() {
	fmt.Println("Hi Pawan Kumar :")
	time.Sleep(3 * time.Second)
}

func main() {
	fmt.Println("Learning goroutine")
	go sayHello()
	go sayHi()
	time.Sleep(5 * time.Second)
}
