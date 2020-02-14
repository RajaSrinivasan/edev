package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the history of interactions",
	Long: `

	Print a list of all files available since the mod time of the specified file.
	For device nodes, only files designated for the node will be listed.

	For the node registered as the publisher node, print a list of all upload/download activity.
	If arguments are specified, then list only for those devices. Otherwise activity of all the nodes
	are listed. This will include uploads and downloads.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			log.Printf("Showing activity for %v", args)
		} else {
			log.Printf("Showing activity for all node")
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	newerThan = showCmd.PersistentFlags().StringP("newer-than", "n", "", "Reference file. Mod time is used to get files newer than")
}
