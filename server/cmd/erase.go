package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// eraseCmd represents the erase command
var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "Erase / cleanup of uploads",
	Long: `
	Erase all the uploads of the given device.

	The argument is the name of the device whose artifacts will be removed.
	`,
	Run: Erase,
}

func init() {
	rootCmd.AddCommand(eraseCmd)
}

func Erase(cmd *cobra.Command, args []string) {
	fmt.Println("erase called")
}
