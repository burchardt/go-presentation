package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone_number,omitempty"`
	Password  string `json:"-"`
}

func main() {
	person := Person{FirstName: "James", LastName: "Bond", Phone: "7007007", Password: "admin1"}

	if jsonData, err := json.Marshal(&person); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(jsonData))
	}
}
