/*
Copyright © 2024 Alex Young <>
*/
package cmd

import (
	"os"

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
    // Existing configuration code
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("apiKey", "")
		viper.Set("location", "")
		viper.Set("unit", "metric")
		viper.SafeWriteConfig()
	}
    viper.AutomaticEnv()
}
