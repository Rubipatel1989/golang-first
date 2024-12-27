package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning web request")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World")
	// })
	// log.Fatal(http.ListenAndServe(":8080", nil))
	myUrl := "https://jsonplaceholder.typicode.com/users/1?value=123&age=23"
	fmt.Printf("Type of url : %T\n", myUrl)
	parsedUrl, error := url.Parse(myUrl)
	if error != nil {
		fmt.Println("Error is", error)
		fmt.Println(error)
		return
	}
	fmt.Printf("Type of formedUrl %T\n", parsedUrl)
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)
	parsedUrl.Path = "/newPath"
	parsedUrl.RawQuery = "name=pawan&lastname=kumar"
	newUrl := parsedUrl.String()
	fmt.Println("New Url", newUrl)
}
