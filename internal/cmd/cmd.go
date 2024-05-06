package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Package = "myple"
	Short   = "Work seamlessly with Myple from the command-line"
	Version = "0.1.0"

	// palette
	Ghost      = lipgloss.Color("240")
	Foreground = lipgloss.Color("#F5F5F5")
	Background = lipgloss.Color("#888B7E")
	Primary    = lipgloss.Color("#7980DE")
)

var (
	// global flags
	// env   string
	debug bool
)

var rootCmd = &cobra.Command{
	Use:     Package,
	Short:   Short,
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute root")
		os.Exit(1)
	}
}

func init() {
	// - global flags
	viper.SetEnvPrefix(Package)
	viper.AutomaticEnv()

	// TODO check if there is a way to highlight that this flag as already been set by env var like in clap
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
	err := viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to bind debug flag")
	}

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Output usage information")
	rootCmd.SetVersionTemplate(fmt.Sprintf("%s %s", Package, Version))

	// - cmds
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(whoamiCmd)

	ApplyStyle(rootCmd)
}
