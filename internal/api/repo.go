package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/universtar-org/tools/internal/model"
)

// GetRepo Get repo information including description, number of stars, etc., via GitHub API.
func (c *Client) GetRepo(ctx context.Context, owner, repo string) (*model.Repo, int, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/repos/%s/%s", owner, repo))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var r model.Repo
	status, err := c.do(req, &r)
	if err != nil {
		return nil, status, err
	}

	return &r, status, nil
}
