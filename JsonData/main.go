package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Creating a JSON object
	person := Person{Name: "Alice", Age: 30, Address: "123 Main St"}

	// Encoding the object to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	// Decoding JSON to an object
	var decodedPerson Person
	if err := json.Unmarshal(jsonData, &decodedPerson); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Decoded:", decodedPerson)
}
