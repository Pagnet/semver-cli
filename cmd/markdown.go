package cmd

import (
	"os"

	"github.com/gavv/cobradoc"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	GroupID: groupDocs,
	Use:     "markdown",
	Short:   "Genereate markdown documentation",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		cobradoc.WriteDocument(os.Stdout, rootCmd, cobradoc.Markdown, cobradoc.Options{
			LongDescription: "Example command using cobra and cobradoc.",
		})

	},
}

func init() {
	rootCmd.AddCommand(markdownCmd)

}
