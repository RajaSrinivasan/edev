package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pendingCmd = &cobra.Command{
	Use:   "pending",
	Short: "Show authorizations pending",
	Long: `
	Shows the devices awaiting authorization.`,
	Run: Pending,
}

func init() {
	showCmd.AddCommand(pendingCmd)
}

func Pending(cmd *cobra.Command, args []string) {
	fmt.Println("pending called")
}
