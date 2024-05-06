package cmd

import "github.com/spf13/cobra"

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Performs a deployment",
	Run: func(_ *cobra.Command, _ []string) {
		// ...
	},
}
