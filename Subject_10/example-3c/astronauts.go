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

type WebRequest interface {
	FetchBody(string) ([]byte, error)
}

type ProxyWebRequest struct {
	HttpProxy string
}

func getHttpClient(httpProxy string) (*http.Client, error) {
	proxyURL, err := url.Parse(httpProxy)
	if err != nil {
		return &http.Client{}, err
	}

	transport := http.Transport{Proxy: http.ProxyURL(proxyURL)}
	httpClient := &http.Client{
		Timeout:   3 * time.Second,
		Transport: &transport,
	}
	return httpClient, nil
}

func FetchBody(httpClient *http.Client, addr string) ([]byte, error) {
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

func (p ProxyWebRequest) FetchBody(addr string) ([]byte, error) {
	httpClient, err := getHttpClient(p.HttpProxy)
	if err != nil {
		return []byte{}, err
	}

	return FetchBody(httpClient, addr)
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

const proxy = "http://proxy.lbs.alcatel-lucent.com:8000"

func GetAstronauts(webRequest WebRequest, addr string) (*Astronauts, error) {
	bodyBytes, err := webRequest.FetchBody(addr)
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
	const addr = "http://api.open-notify.org/astros.json"
	astros, err := GetAstronauts(ProxyWebRequest{HttpProxy: proxy}, addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Astronauts: %v\n", astros.Crew)
}
