package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"cli/pkg/models"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) Login(user models.User) (bool, error) {
	data, _ := json.Marshal(user)
	resp, err := http.Post(c.BaseURL+"/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}
