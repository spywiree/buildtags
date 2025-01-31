package buildtags

import (
	"runtime/debug"
	"strings"
)

// retrieveBuildTags fetches the build tags from the build information.
func retrieveBuildTags() []string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}

	for _, setting := range info.Settings {
		if setting.Key == "-tags" {
			return strings.Split(setting.Value, ",")
		}
	}
	return nil
}

var tags = retrieveBuildTags()

// Tags returns the list of build tags.
func Tags() []string {
	return tags
}

// Has checks if a specific tag is present in the build tags.
func Has(tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}
