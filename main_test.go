package sethealth

import (
	"fmt"
	"testing"
)

func TestValidRequest(t *testing.T) {
	client := New()
	token, err := client.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	if len(token) < 10 {
		t.Fatal("TOKEN is not valid")
	}
}

func TestUnvalidRequest(t *testing.T) {
	client := NewWithCredentials("", "")
	ta, err := client.GetToken()
	fmt.Println(ta, err)
	if err == nil {
		t.Fatal("Request should be unvalid")
	}
}
