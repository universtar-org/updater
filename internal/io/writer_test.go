package io_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/universtar-org/tools/internal/io"
	"github.com/universtar-org/tools/internal/model"
	"github.com/universtar-org/tools/internal/utils"
)

func TestWriteYaml(t *testing.T) {
	tmp := t.TempDir()
	outPath := filepath.Join(tmp, "projects.yaml")
	want := []model.Project{
		{
			Repo:        "universtar",
			Stars:       0,
			Description: "A platform for university students to showcase and promote their projects.",
			Tags: []string{
				"HTML",
				"hugo",
				"universtar",
			},
			UpdatedAt: "2026-02-03T13:56:19Z",
		},
		{
			Repo:        "tools",
			Stars:       0,
			Description: "A tool used to fetch and update project data for universtar.",
			Tags:        []string{"Go"},
			UpdatedAt:   "2026-02-03T13:58:00Z",
		},
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		projects []model.Project
		path     string
		wantErr  bool
	}{
		{
			name:     "WriteYAML",
			projects: want,
			path:     outPath,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := io.WriteYaml(tt.projects, tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("WriteYaml() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("WriteYaml() succeeded unexpectedly")
			}
			data, _ := os.ReadFile(outPath)
			var got []model.Project
			yaml.Unmarshal(data, &got)
			utils.HandleTestDiff(t, want, got)
		})
	}
}
