package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downloadHistory *string

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Download files",
	Long: `
	Download files that are allowed for my device.

	If an argument is provided then only the specified name will be downloaded.

	Otherwise:
		A history file is provided to ascertain the last download. 
		Only newer downloads are sent. Only 1 file is sent for each invocation.

	All files are .spm files.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	downloadHistory = pushCmd.PersistentFlags().StringP("history-file", "H", "", "Download history file")

}
