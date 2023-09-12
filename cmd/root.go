package cmd

import (
	"fmt"
	"os"

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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing CLI '%s'", err)
		os.Exit(1)
	}
}
