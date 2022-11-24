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
	var person Person
	jsonData := []byte(`{"first_name":"James", "last_name":"Bond", "phone_number":"7007007"}`)

	if err := json.Unmarshal(jsonData, &person); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(fmt.Sprintf("%#v", person))
	}
}
