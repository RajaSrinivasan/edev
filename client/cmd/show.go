package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var since *string

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the history of interactions",
	Long: `

	Print a list of all upload/download activity for the specified
	device.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	since = showCmd.PersistentFlags().StringP("since-date", "S", "", "History since the provided date")
}
