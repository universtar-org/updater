package api

import (
	"context"
	"fmt"
	"net/http"
)

type Repo struct {
	Description string   `json:"description"`
	Stars       int      `json:"stargazers_count"`
	Tags        []string `json:"topics"`
	Language    string   `json:"language"`
	UpdatedAt   string   `json:"updated_at"`
}

// GetRepo Get repo information including description, number of stars, etc., via GitHub API.
func (c *Client) GetRepo(ctx context.Context, owner, repo string) (*Repo, int, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/repos/%s/%s", owner, repo))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var r Repo
	status, err := c.do(req, &r)
	if err != nil {
		return nil, status, err
	}

	return &r, status, nil
}
