package get

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"alexleyoung/sun/types"
	"alexleyoung/sun/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get weather information",
	Long: `Get weather information for the current location or for an arbitrary location.`, 
	Run: getDefault,
}

func getDefault(cmd *cobra.Command, args []string) {
	// if no api key set, panic
	location := viper.Get("location")
	apiKey := viper.Get("apiKey")
	apiKeyStr, ok1 := apiKey.(string)
    locationStr, ok2 := location.(string)

    if !ok1 || !ok2 {
        panic("apiKey and location must both be strings")
    }

	body := utils.GetWeatherRestOfDay(apiKeyStr, locationStr)

	var weather types.Weather

	err := json.Unmarshal(body, &weather)
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

