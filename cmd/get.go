package cmd

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get weather information",
	Long: `Get weather information for the current location or for an arbitrary location.`, 
	Run: getWeather,
}

func getWeather(cmd *cobra.Command, args []string) {
}

var location string

func init() {
	getCmd.Flags().StringVarP(&location, "location", "l", "", "Location to get weather information for")
	rootCmd.AddCommand(getCmd)
}