package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	token   string
	http    *http.Client
}

// NewClient Create a `client` with provided `token` which can be empty if the API does not need a token in the header.
func NewClient(token string) *Client {
	return &Client{
		baseURL: "https://api.github.com",
		token:   token,
		http:    &http.Client{Timeout: 10 * time.Second},
	}
}

// newRequest Create a new request before sending.
func (c *Client) newRequest(ctx context.Context, method, path string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	return req, nil
}

// do Send the request and decode the response in `json` format.
func (c *Client) do(req *http.Request, v any) (int, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return http.StatusBadRequest, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return resp.StatusCode, fmt.Errorf("github api error: %s", resp.Status)
	}

	if v != nil {
		return resp.StatusCode, json.NewDecoder(resp.Body).Decode(v)
	}

	return resp.StatusCode, nil
}
