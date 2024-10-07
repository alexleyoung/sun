package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	if viper.Get("apiKey") == "" || viper.GetString("location") == "" {
		fmt.Println("Error: No API key or Location set. Please run `sun config set` to set your API key and location.")
		return
	}

	location := viper.Get("location")
	apiKey := viper.Get("apiKey")

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s", apiKey, location))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Error: Request failed with status code", res.StatusCode)
		return
	}

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

	// Print current weather
	fmt.Println("Current Weather:")
	fmt.Printf("Low: %.1f°F, High: %.1f°F\n\n", weather.Forecast.Forecastday[0].Day.MinTempF, weather.Forecast.Forecastday[0].Day.MaxTempF)

	// Print today's forecast
	fmt.Println("Todays Forecast:")
	for hour := range weather.Forecast.Forecastday[0].Hour {
		hourInfo := weather.Forecast.Forecastday[0].Hour[hour]
		location, err := time.LoadLocation(weather.Location.TzID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Get the current hour
		currentHour, err := strconv.ParseInt((time.Now().In(location).Format("15:04")[0:2]), 10, 0)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// If the current hour is greater than the hour we are iterating through, skip it
		if currentHour > int64(hour) {
			continue
		}

		// Format the time
		hourTime := strings.Split(hourInfo.Time, " ")[1]
		hourTimeObj, err := time.Parse("15:04", hourTime)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		// Format the forecast message
		message := fmt.Sprintf("%s: %.1f°F", hourTimeObj.Format("3:04 PM"), hourInfo.TempF)
		if hourInfo.WillItRain == 1 {
			message += " It will rain!\n"
		} else if hourInfo.WillItSnow == 1 {
			message += " It will snow!\n"
		}
		fmt.Println(message)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
}