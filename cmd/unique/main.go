package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/universtar-org/tools/internal/api"
	"github.com/universtar-org/tools/internal/model"
	"github.com/universtar-org/tools/internal/utils"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("Usage: unique ${username}"))
	}

	client, ctx := utils.InitClientAndContext("")
	username := os.Args[1]

	user, err := client.GetUser(ctx, username)
	if err != nil {
		panic(err)
	}

	if username != user.Name {
		panic(fmt.Errorf("Got: %s\nExpect: %s\n", username, user.Name))
	}

	repos, status, err := client.GetRepoByUser(ctx, username)
	if err != nil {
		panic(err)
	}
	if status != http.StatusOK {
		panic(fmt.Errorf("Get %v", status))
	}

	checkUniqueness(client, ctx, repos, *user)
}

func checkUniqueness(client *api.Client, ctx context.Context, repos []model.Repo, user model.User) {
	if user.Type != "User" {
		return
	}

	for _, repo := range repos {
		if repo.Name == "tools" {
			continue
		}
		contents, err := client.GetDirContent(ctx, "universtar-org", repo.Name, "data/projects")
		if err != nil {
			panic(err)
		}

		for _, content := range contents {
			if user.Name == strings.TrimSuffix(content, filepath.Ext(content)) {
				panic(fmt.Errorf("Duplicated Username in universtar-org/%s", repo.Name))
			}
		}
	}
}
