package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Long: `Set a configuration value. Expects a key and value.`,
	Example: `sun config set apiKey 1234567890abcdef`,
	Run: set,
}

func set(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Error: Not enough arguments. Usage: set <key> <value>")
		return
	}

	key := args[0]
	value := args[1]
	
	// Check if the key already exists
	if viper.Get(key) == nil {
		fmt.Println("Error: Key does not exist")
		return
	}

	// Check if the value for unit is valid
	if viper.Get(key) == "unit" {
		if value != "metric" && value != "imperial" {
			fmt.Println("Error: Invalid unit. Must be either metric or imperial")
			return
		}
	}

	viper.Set(key, value)
	viper.WriteConfig()
}

func init() {
	ConfigCmd.AddCommand(setCmd)
}