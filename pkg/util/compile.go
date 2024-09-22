package util

import (
	"go/build"
	"strings"
)

func GetBuildFlags() []string {
	if len(build.Default.BuildTags) == 0 {
		return nil
	}
	return []string{"-tags=" + strings.Join(build.Default.BuildTags, ",")}
}
