package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	groupMain = "main"
	groupDocs = "docs"
)

var rootCmd = &cobra.Command{
	Use:   "semver-cli",
	Short: "CLI tool to manage semver",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddGroup(&cobra.Group{
		Title: "Main Commands",
		ID:    groupMain,
	})
	rootCmd.AddGroup(&cobra.Group{
		Title: "Documentation Commands",
		ID:    groupDocs,
	})
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
