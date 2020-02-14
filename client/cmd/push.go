package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var targetnodes *[]string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push files to the server.",
	Long: `
	One or more files are pushed to the server.

	Admin nodes push artifacts for distribution to specified nodes.
	Device nodes push logs and other similar data files to the server.

	Only spm files can be pushed.`,
	Args: cobra.MinimumNArgs(1),
	Run:  Push,
}

func init() {

	rootCmd.AddCommand(pushCmd)
	targetnodes = pushCmd.PersistentFlags().StringArrayP("target-nodes", "t", []string{""}, "Target nodes")

}

func Push(cmd *cobra.Command, args []string) {
	log.Printf("Will push %s to %v", args[0], *targetnodes)
}
