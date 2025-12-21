package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	go func(val int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	// time.Sleep(1 * time.Second)
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
