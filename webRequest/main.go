package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//fmt.Println("Learning web request")
	response, error := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if error != nil {
		fmt.Println("Error is", error)
		fmt.Println(error)
		return
	}
	defer response.Body.Close()
	fmt.Println("Response data type is", fmt.Sprintf("%T", response))
	finalDta, _ := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("Error is", error)
		fmt.Println(error)
		return
	}
	fmt.Println("Final data is", string(finalDta))
}
