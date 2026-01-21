package main

func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	// Example usage
	original := "ATAL"
	//original := "1,2,3,4,5"
	reversed := Reverse(original)
	println("Original:", original)
	println("Reversed:", reversed)
}
