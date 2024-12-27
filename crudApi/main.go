package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning Crud operation start")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Println("Get Error is", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting in response code", res.Status)
		return
	}
	var toDo Todo
	err = json.NewDecoder(res.Body).Decode(&toDo)
	if err != nil {
		fmt.Println("Error decoding", res)
		return
	}
	fmt.Println("Get API Response", toDo)
	// data, error := ioutil.ReadAll(res.Body)
	// if error != nil {
	// 	fmt.Println("Get Error Conversion is", err)
	// 	return
	// }
	// fmt.Println("Get API Response", string(data))

}
