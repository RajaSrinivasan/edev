package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the server",
	Long: `
	Installation requires the setup of the backend including creating the
	self signed certificates for the service. `,
	Run: Install,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func Install(cmd *cobra.Command, args []string) {
	fmt.Println("install called")
}
