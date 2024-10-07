package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"alexleyoung/sun/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get weather information",
	Long: `Get weather information for the current location or for an arbitrary location.`, 
	Run: getWeather,
}

func getWeather(cmd *cobra.Command, args []string) {
	// if no api key set, prompt user for api key
	if viper.Get("apiKey") == "" {
		fmt.Print("Enter API key: ")
		var apiKey string
		fmt.Scanln(&apiKey)
		viper.Set("apiKey", apiKey)
	}

	// if no location set, prompt user for location
	if viper.Get("location") == "" {
		fmt.Print("Enter Location (zip OR city OR state OR country): ")
		var location string
		fmt.Scanln(&location)
		viper.Set("location", location)
	} 

	location := viper.Get("location")
	apiKey := viper.Get("apiKey")

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s", apiKey, location))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Status:", res.Status)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var weather types.Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(weather.Current.TempC)
}

func init() {
	rootCmd.AddCommand(getCmd)
}