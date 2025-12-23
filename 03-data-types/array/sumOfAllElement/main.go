package main

func main() {
	num := [5]int{10, 20, 30, 40, 50}
	sum := 0
	// for _, value := range num {
	// 	sum += value
	// }
	for _, value := range num {
		sum += value
	}
	println("Sum of all elements in array is ", sum)
}
