/*
Copyright © 2024 Alex Young <>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sun",
	Short: "CLI to quickly get weather information",
	Long: 
`Sun is a CLI to quickly get weather information.
It can be used to quickly get weather information for a specific location or for the current location. 

The weather information is retrieved from the WeatherAPI API. 
The API key is stored in the config file. 
If the API key is not found, the user will be prompted to enter it.`,
	Example: `# Get weather information for a specific location
  sun -l "New York, NY"

  # Get weather information for the current location
  sun -c

  # Get weather information for a specific location and display the results in JSON format
  sun -l "New York, NY" -j

  # Get weather information for the current location and display the results in JSON format
  sun -c -j`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sun.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


