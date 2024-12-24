package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("We are learning JSON")
	person := Person{Name: "Pawan Kumar", Age: 30, Address: "2/148 Vinamra Khand Gomti Nagar", IsAdult: true}
	fmt.Println("Persion Detauls", person)

	// convert persion into JSON Encoding [Marshling]
	jsonData, error := json.Marshal(person)
	if error != nil {
		fmt.Println("Error is: ", error)
		return
	}
	fmt.Println("Persion Data is ", string(jsonData))

	// Decoding  [Unmarsling]
	var personData Person
	error = json.Unmarshal(jsonData, &personData)
	if error != nil {
		fmt.Println("Error unmarsling : ", error)
		return
	}
	fmt.Println("Persion data is after unmarshall", personData)

}
