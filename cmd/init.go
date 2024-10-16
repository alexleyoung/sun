package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the sun CLI",
	Long: `Initialize the sun CLI. This will prompt you for your API key and location. 
	Get an API key from the weatherapi.com website.  
	The location can be a City or Zipcode`, 
	Run: startInitFlow,
}

func startInitFlow(cmd *cobra.Command, args []string) {
	// Prompt for API key
	var apiKey string
	fmt.Print("Enter your API key: ")
	fmt.Scanln(&apiKey)

	// Prompt for location
	var location string
	fmt.Print("Enter your location: ")
	fmt.Scanln(&location)

	// Set the API key and location in the config
	viper.Set("apiKey", apiKey)
	viper.Set("location", location)
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}

	// Print a success message
	println("Successfully initialized the sun CLI!")
}

func init() {
	rootCmd.AddCommand(initCmd)
}