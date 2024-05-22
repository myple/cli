package cmd

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logs into your account or creates a new one",
	Run: func(_ *cobra.Command, _ []string) {
		err := open.Run("https://app.myple.io/login")
		if err != nil {
			fmt.Println("failed opening browser. Copy the url (https://app.myple.io/login) into a browser and continue")
		}
	},
}
