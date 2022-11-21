package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:8002/hello")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
