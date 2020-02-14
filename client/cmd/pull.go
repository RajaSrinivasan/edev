package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var newerThan *string

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Download files",
	Long: `
	Download files that are allowed for my device.

	If an argument is provided then only the specified name will be downloaded.

	All files are .spm files.

	`,
	Run: Pull,
}

func init() {
	rootCmd.AddCommand(pullCmd)
	newerThan = pullCmd.PersistentFlags().StringP("newer-than", "n", "", "Reference file. Mod time is used to get files newer than")
}

func Pull(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		log.Printf("Will pull the file %s", args[0])
	} else {
		log.Println("Will pull all files available")
	}
}
