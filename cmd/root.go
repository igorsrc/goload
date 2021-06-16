package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "goload",
	Version: "v1.0",
	Short:   "short description",
	Long:    "Goload v1.0",
}

func init() {
	rootCmd.AddCommand(cmdGet)
	rootCmd.AddCommand(cmdPost)
	rootCmd.AddCommand(cmdPut)
}

func Execute() error {
	return rootCmd.Execute()
}
