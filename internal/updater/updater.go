package updater

import (
	"context"
	"fmt"
	"net/http"

	"github.com/universtar-org/tools/internal/api"
	"github.com/universtar-org/tools/internal/io"
	"github.com/universtar-org/tools/internal/utils"
)

// Update a file/user.
func Update(client *api.Client, ctx context.Context, path string) error {
	const MAX_TAG_NUMBER = 5
	owner := utils.ParseOwner(path)

	projects, err := io.ReadYaml(path)
	if err != nil {
		return err
	}

	for i := range projects {
		fmt.Printf("Processing: %v/%v\n", owner, projects[i].Repo)

		repo, status, err := client.GetRepo(ctx, owner, projects[i].Repo)
		if err != nil {
			return err
		}

		if status != http.StatusOK {
			return fmt.Errorf("Request Failed: Get %v", status)
		}

		tag_list := append([]string{repo.Language}, repo.Tags...)
		if len(tag_list) > MAX_TAG_NUMBER {
			tag_list = tag_list[:MAX_TAG_NUMBER]
		}
		projects[i].Description = repo.Description
		projects[i].Stars = repo.Stars
		projects[i].UpdatedAt = repo.UpdatedAt
		projects[i].Tags = tag_list
	}

	io.WriteYaml(projects, path)

	return nil
}
