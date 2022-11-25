package main

import (
	"testing"
)

type TestWebRequest struct {
}

func (TestWebRequest) FetchBody(string) ([]byte, error) {
	return []byte(`{"message": "OK", "number": 2, "people": [{"name": "Neil Armstrong", "crew": "NASA"}, {"name": "Pete Conrad", "crew": "NASA"}]}`), nil
}

func TestGetAstronauts(t *testing.T) {
	astros, err := getAstronauts(TestWebRequest{}, "")
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
