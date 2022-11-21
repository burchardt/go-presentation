package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Users struct {
	Total int `json:"total"`
	Data  []struct {
		id        int    `json:"id"`
		email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"data"`
	Support struct {
		URL  string `json:"id"`
		Text string `json:"text"`
	} `json:"support"`
}

func main() {
	os.Setenv("HTTP_PROXY", "http://proxy.lbs.alcatel-lucent.com:8000")
	response, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var users Users
	if err := json.Unmarshal(body, &users); err != nil {
		log.Fatal(err)
	}
	for _, user := range users.Data {
		fmt.Printf("%s %s\n", user.FirstName, user.LastName)
	}
}
