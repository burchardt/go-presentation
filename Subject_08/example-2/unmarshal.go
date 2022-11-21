package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone_number,omitempty"`
	Password  string `json:"-"`
}

func main() {
	var person Person
	jsonData := []byte(`{"first_name":"Janusz","last_name":"Kowalski","phone_number":"1234567"}`)

	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}
}
