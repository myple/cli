package cmd

import "github.com/spf13/cobra"

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Shows the username of the currently logged in user",
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}
