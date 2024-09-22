package util

import (
	"go/build"
	"testing"
)

func TestGetBuildFlags(t *testing.T) {
	build.Default.BuildTags = []string{"test"}
	buildFlags := GetBuildFlags()
	if len(buildFlags) != 1 {
		t.Errorf("Expected 1 build flag, got %d", len(buildFlags))
	}
	if buildFlags[0] != "-tags=test" {
		t.Errorf("Expected build flag to be '-tags=test', got '%s'", buildFlags[0])
	}
}
