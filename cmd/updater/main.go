package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/universtar-org/tools/internal/api"
	"github.com/universtar-org/tools/internal/io"
	"github.com/universtar-org/tools/internal/utils"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("Usage: updater ${data-file-dir}"))
	}
	client, ctx := utils.InitClientAndContext("")

	list, err := io.GetDataFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, path := range list {
		fmt.Println("Processing: ", path)
		if err := update(client, ctx, path); err != nil {
			panic(err)
		}
	}

	fmt.Println("Finished!")
}

// Update a file/user.
func update(client *api.Client, ctx context.Context, path string) error {
	const maxTagNumber = 5
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

		tags := append([]string{repo.Language}, repo.Tags...)
		if len(tags) > maxTagNumber {
			tags = tags[:maxTagNumber]
		}
		projects[i].Description = repo.Description
		projects[i].Stars = repo.Stars
		projects[i].UpdatedAt = repo.UpdatedAt
		projects[i].Tags = tags
	}

	io.WriteYaml(projects, path)

	return nil
}
