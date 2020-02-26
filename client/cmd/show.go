package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/edev/client/impl"
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
	Run: Show,
}

func init() {
	rootCmd.AddCommand(showCmd)
	newerThan = showCmd.PersistentFlags().StringP("newer-than", "n", "", "Reference file. Mod time is used to get files newer than")
}

func Show(cmd *cobra.Command, args []string) {
	arg := "all"
	if len(args) > 0 {
		arg = args[0]
	}
	impl.Show(arg)

}
