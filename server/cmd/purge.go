package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var retentionPeriod int

// showCmd represents the show command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge artifacts",
	Long: `
	Device artifacts (logs etc) have a retention policy.
	Purge command erases files that are older.`,
	Run: Purge,
}

func init() {
	rootCmd.AddCommand(purgeCmd)
	purgeCmd.PersistentFlags().IntVarP(&retentionPeriod, "retention-period", "r", 30, "Retention Period in no of days. default 30")
}

func Purge(cmd *cobra.Command, args []string) {
	fmt.Println("Purge called")
}
