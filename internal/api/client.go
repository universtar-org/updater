package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
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
		http: &http.Client{
			Timeout:       10 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		},
	}
}

// newRequest Create a new request before sending.
func (c *Client) newRequest(ctx context.Context, method, path string) (*http.Request, error) {
	slog.Debug(
		"build request",
		"method", method,
		"path", path,
	)

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("create request %s %s: %w", method, path, err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	return req, nil
}

// do Send the request and decode the response in `json` format.
func (c *Client) do(req *http.Request, v any) (int, error) {
	slog.Debug("send request",
		"method", req.Method,
		"url", req.URL.String(),
	)

	const maxRetries = 5
	var resp *http.Response
	var err error

	for i := range maxRetries {
		resp, err = c.http.Do(req)
		if err != nil {
			slog.Info("send request failed",
				"times", i+1,
				"err", err,
			)
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			slog.Info(
				"send request failed",
				"times", i+1,
				"status", resp.Status,
			)
			resp.Body.Close()
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}
	}

	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("send request %s %s: %w", req.Method, req.URL.String(), err)
	}

	defer resp.Body.Close()
	slog.Debug(
		"receive response",
		"method", req.Method,
		"url", req.URL.String(),
		"status", resp.StatusCode,
	)

	if resp.StatusCode >= 400 {
		return resp.StatusCode, fmt.Errorf("github api error: %s", resp.Status)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return resp.StatusCode, fmt.Errorf("decode response: %w", err)
		}
	}

	return resp.StatusCode, nil
}
