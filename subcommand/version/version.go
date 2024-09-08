package version

import (
	"fmt"
	"runtime/debug"
)

var version string = "devel"

type Version struct {
}

func (Version) Execute(arguments []string) (string, error) {
	result := ""
	info, _ := debug.ReadBuildInfo()
	result += fmt.Sprintf("LINE-Webhook-Karaoke %s\n", version)
	result += fmt.Sprintf("%s\n", info.GoVersion)
	for _, dep := range info.Deps {
		result += fmt.Sprintf("%s %s\n", dep.Path, dep.Version)
	}
	return result, nil
}

func (Version) Name() string {
	return "version"
}
