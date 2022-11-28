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

func FetchBody(addr string) ([]byte, error) {
	const proxy = "http://proxy.lbs.alcatel-lucent.com:8000"

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return []byte{}, err
	}

	transport := http.Transport{Proxy: http.ProxyURL(proxyURL)}
	httpClient := http.Client{
		Timeout:   3 * time.Second,
		Transport: &transport,
	}
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return []byte{}, err
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return []byte{}, err
	}
	return body, nil
}

type Member struct {
	Name  string `json:"name"`
	Cratf string `json:"craft"`
}
type Astronauts struct {
	Number  int      `json:"number"`
	Message string   `json:"message"`
	Crew    []Member `json:"people"`
}

func GetAstronauts() (*Astronauts, error) {
	const addr = "http://api.open-notify.org/astros.json"

	bodyBytes, err := FetchBody(addr)
	if err != nil {
		return &Astronauts{}, err
	}
	astronauts := Astronauts{}
	err = json.Unmarshal(bodyBytes, &astronauts)
	if err != nil {
		return &Astronauts{}, err
	}
	return &astronauts, nil
}

func main() {
	astros, err := GetAstronauts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Astronauts: %v\n", astros.Crew)
}
