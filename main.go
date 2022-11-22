package main

import (
	_ "embed"

	"github.com/Pagnet/semver-cli/cmd"
)

var (
	version string
)

func main() {
	cmd.Version = version
	cmd.Execute()
}
