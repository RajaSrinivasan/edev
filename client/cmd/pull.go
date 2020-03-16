package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var newerThan *string
var listOption *bool
var outputDir *string
var devNames *[]string

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Retrieve list of available files",
	Long: `
	Retrieve a list of files that are available for my device. The returned list is a list of
	file id's.

	Files can be retrieved using file id's as arguments

	`,

	Run: Pull,
}

func init() {
	rootCmd.AddCommand(pullCmd)
	newerThan = pullCmd.PersistentFlags().StringP("since", "s", "", "Reference file. Mod time is used to get files newer than. Only list these files")
	listOption = pullCmd.PersistentFlags().BoolP("list-only", "l", false, "Retrieve a list of the files.")
	devNames = pullCmd.PersistentFlags().StringArrayP("device-names", "d", []string{""}, "Devices Names whose files are to be listed")
	outputDir = pullCmd.PersistentFlags().StringP("output-dir", "o", "", "Output Dir where files are to be stored")
}

func Pull(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		log.Printf("Will pull the file %s", args[0])
	} else {
		log.Println("Will pull all files available")
	}
}
