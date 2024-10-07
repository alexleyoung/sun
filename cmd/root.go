/*
Copyright © 2024 Alex Young <>
*/
package cmd

import (
	"fmt"
	"os"

	"alexleyoung/sun/cmd/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sun",
	Short: "CLI to quickly get weather information",
	Long: 
`Sun is a CLI to quickly get weather information.
It can be used to quickly get weather information for a specific location or for the current location. 
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

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
    // Existing configuration code
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("No config file found, using defaults")
    }
}
