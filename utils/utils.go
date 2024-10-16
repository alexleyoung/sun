package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitConfig() {
	// if config already exists, read it
	if err := viper.ReadInConfig(); err == nil {
		return
	}

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

	if err = viper.ReadInConfig(); err != nil {
		// Set default values if the config file does not exist
		viper.Set("apiKey", "")
		viper.Set("location", "")
		viper.Set("unit", "imperial")

		// Now write the config to the specified file path
		if err := viper.WriteConfigAs(configFile); err != nil {
			fmt.Println("Error writing config file:", err)
			return
		}
	}
}