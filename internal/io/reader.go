package io

import (
	"github.com/goccy/go-yaml"
	"github.com/universtar-org/updater/internal/model"

	"os"
	"path/filepath"
)

func ReadYaml(path string) ([]model.Project, error) {
	var projects []model.Project
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func GetDataFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, entry := range entries {
		if !entry.IsDir() {
			paths = append(paths, filepath.Join(dir, entry.Name()))
		}
	}

	return paths, nil
}
