package main

// Assertions are way to extract underline concreteType interface value.
// Always use ok [with  comma ] idiom to avoid panic at runtime.
func main() {
	var a interface{} = "hello go lang"

	//type assertion
	str, ok := a.(string)
	if ok {
		println("String value:", str)
	} else {
		println("a is not a string")
	}

	//type assertion to int
	num, ok := a.(int)
	if ok {
		println("Integer value:", num)
	} else {
		println("a is not an integer")
	}

	//panic on failed assertion
	// Uncommenting the following lines will cause a panic
	num2 := a.(int)
	println("Integer value:", num2)
}
