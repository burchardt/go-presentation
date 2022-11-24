package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	const addr = "http://localhost:8002/hello"

	httpClient := http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
