package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send result to channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 10, 15}

	// Create a channel [The pipe]
	c := make(chan int)
	fmt.Println("Channel created", c)

	//2. Split the work and start two workers as goroutines
	go sum(s[:len(s)/2], c) // worker 1 for first half
	go sum(s[len(s)/2:], c) // worker 2 for second half

	//3. Receive results from the channel[Block untill data is available]
	x := <-c // receive result from worker 1 [Or 2, depending on which one finishes first]
	y := <-c // receive result from worker 2 [Or 1, depending on which one finishes first]

	fmt.Println(x, y, x+y)

}
