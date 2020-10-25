package cmd

import (
	"github.com/spf13/cobra"
	"github.com/btisdall/dotrr/v2/app/util"

)

var rootCmd = &cobra.Command{
	Use:   "dotrr",
	Short: "dotrr is a program for resolving secrets in dotenv files",
}

// Execute is the main entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		util.Er(err)
	}
}
