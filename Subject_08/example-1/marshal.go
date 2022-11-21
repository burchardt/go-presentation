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
	person := Person{FirstName: "Janusz", LastName: "Kowalski", Phone: "1234567", Password: "admin1"}

	jsonData, err := json.Marshal(&person)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonData))
	}
}
