package sethealth

import (
	"testing"
)

func TestValidRequest(t *testing.T) {
	client := New()
	token, err := client.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	if len(token.Token) < 10 {
		t.Fatal("TOKEN is not valid")
	}
}

func TestUnvalidRequest(t *testing.T) {
	client := NewWithCredentials("", "")
	_, err := client.GetToken()
	if err == nil {
		t.Fatal("Request should be unvalid")
	}
}
