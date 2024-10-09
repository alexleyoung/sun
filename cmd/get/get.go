package get

import (
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get weather information",
	Long: `Get weather information for the current location or for an arbitrary location.`, 
}