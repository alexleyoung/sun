package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration value",
	Long: `Set a configuration value.`,
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

	viper.Set(key, value)
	viper.WriteConfig()
}

func init() {
	ConfigCmd.AddCommand(setCmd)
}