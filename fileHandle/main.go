package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating", err)
		return
	} else {
		fmt.Println("File created successfully")
	}
	content := "Hello Pawan, You are the best"
	bytes, errors := io.WriteString(file, content+"\n\n\n")
	if errors != nil {
		fmt.Println("Error while eritting in file", errors)
		return
	} else {
		fmt.Println("Successfully, added lines", bytes)
	}
	defer file.Close()
}
