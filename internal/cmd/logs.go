package cmd

import "github.com/spf13/cobra"

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Displays the logs for a deployment",
	Run: func(_ *cobra.Command, _ []string) {
		// ...
	},
}
