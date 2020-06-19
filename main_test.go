package sethealth

import (
	"fmt"
	"os"
	"testing"
)

func TestValidRequest(t *testing.T) {
	client := New(os.Getenv("SETHEALTH_KEY"), os.Getenv("SETHEALTH_SECRET"))
	token, err := client.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	if len(token) < 10 {
		t.Fatal("TOKEN is not valid")
	}
}

func TestUnvalidRequest(t *testing.T) {
	client := New("", "")
	ta, err := client.GetToken()
	fmt.Println(ta, err)
	if err == nil {
		t.Fatal("Request should be unvalid")
	}
}
