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
		panic(fmt.Errorf("Usage: checker ${data-file-dir}"))
	}

	client := api.NewClient("")
	ctx := context.Background()

	list, err := io.GetDataFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, path := range list {
		fmt.Println("Checking: ", path)
		check(client, ctx, path)
	}

	fmt.Println("Finished!")
}

func check(client *api.Client, ctx context.Context, path string) {
	projects, err := io.ReadYaml(path)
	owner := utils.ParseOwner(path)
	if err != nil {
		panic(err)
	}

	for _, project := range projects {
		fmt.Printf("Checking: %v/%v\n", owner, project.Repo)

		_, status, err := client.GetRepo(ctx, owner, project.Repo)
		if err != nil || status != http.StatusOK {
			var message string
			if err != nil {
				message = err.Error()
			} else {
				message = fmt.Sprintf("Get %v", status)
			}
			panic(fmt.Errorf("%v/%v: %v", owner, project.Repo, message))
		}
	}
}
