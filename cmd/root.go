/*
Copyright © 2024 Alex Young <>
*/
package cmd

import (
	"os"

	"alexleyoung/sun/cmd/config"
	"alexleyoung/sun/cmd/get"
	"alexleyoung/sun/utils"

	"github.com/spf13/cobra"
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
	utils.InitConfig()
	
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(get.GetCmd)
}

