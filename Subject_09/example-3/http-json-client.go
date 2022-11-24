package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func fetchData(endpoint string) ([]byte, error) {
	const httpProxy = "http://proxy.lbs.alcatel-lucent.com:8000"

	proxyURL, err := url.Parse(httpProxy)
	if err != nil {
		return []byte{}, err
	}

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	httpClient := http.Client{
		Transport: transport,
		Timeout:   3 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return []byte{}, err
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

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
	body, err := fetchData("https://reqres.in/api/users?page=2")
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
