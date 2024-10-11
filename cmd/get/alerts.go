package get

import (
	"fmt"

	"alexleyoung/sun/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var alertsCommand = &cobra.Command{
	Use:   "alerts",
	Short: "Get weather alerts",
	Long: `Get weather alerts for the current location or for an arbitrary location.`, 
	Run: getAlerts,
}

func getAlerts(cmd *cobra.Command, args []string) {
	apiKey, ok1 := viper.Get("apiKey").(string)
	if !ok1 {
		panic("apiKey must be a string")
	} else if apiKey == "" {
		fmt.Println("Set an API key using sun config set.")
		return
	}
	if location == "" {
		fmt.Print("Set a default location or use the -l flag to specify a location.")
		return
	}
	
	alerts := utils.GetAlerts(apiKey, location)

	if len(alerts.Alerts.Alert) == 0 {
		fmt.Println("No current alerts.")
		return
	}

	for _, alert := range alerts.Alerts.Alert {
		color.Red(alert.Headline)
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

	alertsCommand.Flags().StringVarP(&location, "location", "l", locStr, "Location to get alerts for")

	GetCmd.AddCommand(alertsCommand)
}