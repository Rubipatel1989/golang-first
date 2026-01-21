package main

import (
	"fmt"
	"time"
)

func printNumbers(n int) {
	for i := 1; i <= 10; i++ {
		result := n * i
		fmt.Println(n, "*", i, "=", result)
	}
}

func main() {
	println("=========== Start table 1 ============")
	go printNumbers(3)
	time.Sleep(2 * time.Second)
	println("=========== Start table 2 ============")
	printNumbers(2)
}
