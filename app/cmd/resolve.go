package cmd

import (
	"fmt"
	"github.com/btisdall/dottr/v2/app/config"
	"github.com/btisdall/dottr/v2/app/secrets"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(resolveCmd)
}

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve secrets from a dotenv template file",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := config.Read(args[0])
		if err != nil {
			fmt.Printf("Error reading dotenv file: %s\n", err)
			os.Exit(1)
		}
		resolved := Resolve(&content, secrets.GetProvider)
		fmt.Printf("%v\n", config.Marshal(resolved))
	},
}

// Resolve resolves secrets from dotenv file values using the appropriate provider
func Resolve(c *config.Map, getProvider secrets.GetProviderFunction) config.Map {
	channel := make(chan config.Item)

	for k, v := range *c {
		go func(key, value string, c chan config.Item) {
			provider := getProvider(value)
			secret, err := provider.GetSecret(value)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				c <- config.NewItem(key, value)
			} else {
				c <- config.NewItem(key, secret)
			}
		}(k, v, channel)
	}

	resolved := config.Map{}
	for range *c {
		item := <-channel
		resolved[item.Key] = item.Value
	}
	return resolved
}
