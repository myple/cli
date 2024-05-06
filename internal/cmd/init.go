package cmd

import "github.com/spf13/cobra"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an example project",
	Run: func(_ *cobra.Command, _ []string) {
		// ...
	},
}
