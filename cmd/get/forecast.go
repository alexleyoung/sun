package get

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"alexleyoung/sun/utils"

	"github.com/fatih/color"
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
	if days < 1 || days > 14 {
		fmt.Println("Number of days must be between 1 and 14.")
		return
		}
		
	forecast := utils.GetForecast(apiKey, location, days)
	now := time.Now()

	for day := range forecast.Forecast.Forecastday {
		if day == 0 {
			color.HiRed("Today:")
		} else if day == 1 {
			color.HiRed("Tomorrow:")
		} else {
			color.HiRed("Day " + strconv.Itoa(day+1) + ":")
		}

		// Print high and low temperatures
		lowTemp := forecast.Forecast.Forecastday[day].Day.MinTempF
		highTemp := forecast.Forecast.Forecastday[day].Day.MaxTempF
		fmt.Printf("Low: %.1f°F\nHigh: %.1f°F\n\n", lowTemp, highTemp)

		color.Yellow("Forecast:")
		for hour := range forecast.Forecast.Forecastday[day].Hour {
			hourInfo := forecast.Forecast.Forecastday[day].Hour[hour]

			if now.Unix() > hourInfo.TimeEpoch && hour < now.Hour() {
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
				color.Set(color.FgCyan)
				message += " - It will rain!"
			} else if hourInfo.WillItSnow == 1 {
				color.Set(color.FgHiWhite)
				message += " - It will snow!"
			}
			fmt.Println(message)
		}
		fmt.Println()
	}
}

func init() {
	utils.InitConfig()

	viper.ReadInConfig()
	loc := viper.GetString("location")

	forecastCmd.Flags().StringVarP(&location, "location", "l", loc, "Location to get forecast for")
	forecastCmd.Flags().IntVarP(&days, "days", "d", 1, "Number of days to get forecast for")

	GetCmd.AddCommand(forecastCmd)
}