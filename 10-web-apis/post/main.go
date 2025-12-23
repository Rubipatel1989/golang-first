package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
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

}
func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "Pawan Kumar",
		Completed: true,
	}
	jsondata, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("There is error :", err)
		return
	}
	jsonString := string(jsondata)
	jsonReader := strings.NewReader(jsonString)
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("There is error in request :", err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("There is error in response :", err)
		return
	}
	fmt.Println("Response is ", string(data))

}

func main() {
	fmt.Println("Learning Crud operation start")
	//performGetRequest()
	performPostRequest()
}
