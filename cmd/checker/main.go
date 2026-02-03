package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/universtar-org/tools/internal/api"
	"github.com/universtar-org/tools/internal/io"
	"github.com/universtar-org/tools/internal/log"
	"github.com/universtar-org/tools/internal/utils"
)

func main() {
	opts := utils.ParseFlags()
	log.InitLogger(opts.Debug)

	args := flag.Args()
	if len(args) != 1 {
		slog.Error(
			"invalid arguments",
			"usage", "checker <data-file-dir>",
		)
		os.Exit(2)
	}

	client, ctx := utils.InitClientAndContext("")

	dir := args[0]
	list, err := io.GetDataFiles(dir)
	if err != nil {
		slog.Error(
			"failed to get data files",
			"dir", dir,
			"err", err,
		)
		os.Exit(1)
	}

	for _, path := range list {
		slog.Info(
			"checking file",
			"path", path,
		)
		if err := check(client, ctx, path); err != nil {
			slog.Error(
				"check failed",
				"path", path,
				"err", err,
			)
			os.Exit(1)
		}
	}

	slog.Info("finished")
}

func check(client *api.Client, ctx context.Context, path string) error {
	projects, err := io.ReadYaml(path)
	if err != nil {
		return fmt.Errorf("read yaml %s: %w", path, err)
	}

	owner := utils.ParseOwner(path)

	for _, project := range projects {
		slog.Debug("checking repo", "owner", owner, "repo", project.Repo)

		_, status, err := client.GetRepo(ctx, owner, project.Repo)
		if err != nil {
			return fmt.Errorf("checking repo %s/%s: %w", owner, project.Repo, err)
		}

		if status != http.StatusOK {
			return fmt.Errorf("check repo %s/%s: unexpected status %d", owner, project.Repo, status)
		}
	}

	return nil
}
