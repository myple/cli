package cmd

import "github.com/spf13/cobra"

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open documentation in your web browser",
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}
