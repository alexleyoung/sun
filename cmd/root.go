/*
Copyright © 2024 Alex Young <>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"alexleyoung/sun/cmd/config"
	"alexleyoung/sun/cmd/get"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sun",
	Short: "CLI to quickly get weather information",
	Long: 
`Sun is a CLI to quickly get weather information.
Run 'sun init' to get started.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(get.GetCmd)
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Define the config file path
	configDir := filepath.Join(home, ".sun")
	configFile := filepath.Join(configDir, "config.yaml")

	// Create the .sun directory if it doesn't exist
	if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating config directory: %v\n", err)
		return
	}

	// Set Viper configuration path and name
	viper.SetConfigName("config")          // Name of the config file (without extension)
	viper.SetConfigType("yaml")            // Specify config file type
	viper.AddConfigPath(configDir)         // Specify the directory to look for config

	viper.AutomaticEnv()                    // Automatically read env variables that match

	if err := viper.ReadInConfig(); err != nil {
		// Set default values if the config file does not exist
		viper.Set("apiKey", "")
		viper.Set("location", "")
		viper.Set("unit", "imperial")

		// Now write the config to the specified file path
		if err := viper.WriteConfigAs(configFile); err != nil {
			fmt.Println("Error writing config file:", err)
			return
		}
	} else {
		fmt.Println("Config file found:", configFile)
	}
}
