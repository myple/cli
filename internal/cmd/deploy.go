package cmd

import "github.com/spf13/cobra"

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Performs a deployment",
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}
