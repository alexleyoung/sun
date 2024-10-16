package config

import (
	"fmt"

	"alexleyoung/sun/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the sun CLI",
	Long: `Configure the sun CLI. Default prints the current configuration.`,
	Run: config,
}

func config(cmd *cobra.Command, args []string) {
	fmt.Println("API key:", viper.GetString("apiKey"))
	fmt.Println("Location:", viper.GetString("location"))
	fmt.Println("Unit:", viper.GetString("unit"))
}

func init() {
	utils.InitConfig()
}