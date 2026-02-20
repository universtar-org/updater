package io_test

import (
	"github.com/universtar-org/tools/internal/io"
	"github.com/universtar-org/tools/internal/model"
	"github.com/universtar-org/tools/internal/utils"
	"testing"
)

func TestReadYaml(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    []model.Project
		wantErr bool
	}{
		{
			name: "ValidYAML",
			path: "./testdata/valid.yaml",
			want: []model.Project{
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
			},
			wantErr: false,
		},
		{
			name:    "InvalidYAML",
			path:    "./testdata/invalid.yaml",
			wantErr: true,
		},
		{
			name:    "FileNotExist",
			path:    "./testdata/not-exist.yaml",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := io.ReadYaml(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ReadYaml() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ReadYaml() succeeded unexpectedly")
			}
			utils.HandleTestDiff(t, tt.want, got)
		})
	}
}
