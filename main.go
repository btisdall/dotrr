package main

import (
	"fmt"
	"os"

	"github.com/btisdall/dottr/v2/config"
	"github.com/btisdall/dottr/v2/secrets"
)

var (
	version string = "SNAPSHOT"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE: dottr FILE (%s)\n", version)
		os.Exit(1)
	}
	f := os.Args[1]
	content, err := config.Read(f)
	if err != nil {
		fmt.Printf("Error reading dotenv file: %s\n", err)
		os.Exit(1)
	}

	resolved := resolveSecrets(&content, secrets.GetProvider)
	fmt.Printf("%v\n", config.Marshal(resolved))
}

func resolveSecrets(c *config.Map, getProvider secrets.GetProviderFunction) config.Map {
	channel := make(chan config.Item)

	for k, v := range *c {
		go func(key, value string, provider secrets.Provider, c chan config.Item) {
			secret, err := provider.GetSecret(value)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				c <- config.NewItem(key, value)
			} else {
				c <- config.NewItem(key, secret)
			}
		}(k, v, getProvider(v), channel)
	}

	resolved := config.Map{}
	for range *c {
		item := <-channel
		resolved[item.Key] = item.Value
	}
	return resolved
}
