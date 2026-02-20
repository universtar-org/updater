package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func HandleTestDiff(t *testing.T, want, got any) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
