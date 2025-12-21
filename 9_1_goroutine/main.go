package main

import (
	"fmt"
	"time"
)

func printNumber(n int) {
	for i := 0; i <= 10; i++ {
		result := n * i
		//fmt.Println(result)
		fmt.Printf("%d x %d = %d\n", n, i, result)
	}

}
func main() {
	// for i := 0; i <= 10; i++ {
	// 	go func(val int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// time.Sleep(1 * time.Second)
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	go printNumber(2)
	time.Sleep(1 * time.Second)
}
