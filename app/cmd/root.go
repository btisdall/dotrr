package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotrr",
	Short: "dotrr is a program for resolving secrets in dotenv files",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute is the main entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
