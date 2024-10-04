package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the sun CLI",
	Long: `Configure the sun CLI.`,
	Run: config,
}

func config(cmd *cobra.Command, args []string) {
	fmt.Println("Configuration file:", viper.ConfigFileUsed())
	fmt.Println("API key:", viper.GetString("apiKey"))
	fmt.Println("Location:", viper.GetString("location"))
}