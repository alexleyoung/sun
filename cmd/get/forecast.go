package get

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"alexleyoung/sun/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// flags
	location string
	days int

	// cmd
	forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get weather forecast",
	Long: `Get up to 14 days of forecast data for arbitrary location.`, 
	Run: getForecast,
})

func getForecast(cmd *cobra.Command, args []string) {
	apiKey, ok1 := viper.Get("apiKey").(string)
	if !ok1 {
		panic("apiKey must be a string")
	}

	if location == "" {
		fmt.Print("Set a default location or use the -l flag to specify a location.")
		return
	}

	forecast := utils.GetForecast(apiKey, location, days)

	for day := range forecast.Forecast.Forecastday {
		if day == 0 {
			fmt.Println("Today:")
		} else if day == 1 {
			fmt.Println("Tomorrow:")
		} else {
			fmt.Println("Day " + strconv.Itoa(day+1) + ":")
		}
		for hour := range forecast.Forecast.Forecastday[day].Hour {
			hourInfo := forecast.Forecast.Forecastday[day].Hour[hour]
			location, err := time.LoadLocation(forecast.Location.TzID)
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
}

func init() {
	viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
	viper.ReadInConfig()
	loc := viper.Get("location")
	locStr, ok := loc.(string)
	if !ok {
		locStr = ""
	}

	forecastCmd.Flags().StringVarP(&location, "location", "l", locStr, "Location to get forecast for")
	forecastCmd.Flags().IntVarP(&days, "days", "d", 1, "Number of days to get forecast for")

	GetCmd.AddCommand(forecastCmd)
}