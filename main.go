package sethealth

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const host = "https://api.set.health"

// Client exposes the public api for sethealth
type Client struct {
	key    string
	secret string
	client *http.Client
}

type tokenResponse struct {
	Token string `json:"token"`
}

// New creates a new client for the server sethealth API
// It requires a key and secret in order to perform any request
func New(key, secret string) *Client {
	return &Client{
		key:    key,
		secret: secret,
		client: &http.Client{},
	}
}

// GetToken returns a new short-living token to be used by client side.
func (c *Client) GetToken() (string, error) {
	data := map[string]string{
		"key":    c.key,
		"secret": c.secret,
	}
	jsonBytes, _ := json.Marshal(data)
	body := bytes.NewBuffer(jsonBytes)

	// Create request
	req, err := http.NewRequest("POST", host+"/token", body)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	var token tokenResponse
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", err
	}
	return token.Token, nil
}
