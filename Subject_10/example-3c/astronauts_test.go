package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type NoProxyWebRequest struct {
}

func (p NoProxyWebRequest) FetchBody(addr string) ([]byte, error) {
	return FetchBody(&http.Client{}, addr)
}

func TestGetHttpClient(t *testing.T) {
	httpClient, err := getHttpClient(proxy)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %s", err)
	}

	if httpClient.Timeout != 3*time.Second {
		t.Errorf("Timeout, got: %s, want: %s", httpClient.Timeout, 3*time.Second)
	}
}

func TestGetAstronauts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "OK", "number": 2, "people": [{"name": "Neil Armstrong", "crew": "NASA"}, {"name": "Pete Conrad", "crew": "NASA"}]}`))
	}))
	defer server.Close()

	astros, err := GetAstronauts(NoProxyWebRequest{}, server.URL)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %s", err)
	}

	if astros.Number != 2 {
		t.Errorf("Number of people in space, got: %d, want: %d.", astros.Number, 2)
	}

	if astros.Message != "OK" {
		t.Errorf("Message, got: %s, want: %s.", astros.Message, "OK")
	}

	if len(astros.Crew) != 2 {
		t.Errorf("Number of Crew members, got: %d, want: %d.", len(astros.Crew), 2)
	}

}
