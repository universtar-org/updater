package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/universtar-org/tools/internal/model"
)

func (c *Client) GetUser(ctx context.Context, username string) (*model.User, error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/users/%s", username))
	if err != nil {
		return nil, err
	}

	var r model.User
	status, err := c.do(req, &r)
	if err != nil || status != http.StatusOK {
		if err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	}

	return &r, nil
}
