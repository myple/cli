package cmd

import "github.com/spf13/cobra"

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logs out of your account",
	Run: func(_ *cobra.Command, _ []string) {
		// ...
	},
}
