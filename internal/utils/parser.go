package utils

import (
	"path/filepath"
	"strings"
)

func ParseOwner(path string) string {
	base := filepath.Base(path)
	owner := strings.TrimSuffix(base, filepath.Ext(base))

	return owner
}
