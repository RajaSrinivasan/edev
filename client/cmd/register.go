package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nameToregister string
var publisherNode *bool

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register the device",
	Long: `
	Register the device with the server to establish 
	a trusted relationship.	
	Argument is the name to register. In most cases, this will be the same as the hostname.`,
	Args: cobra.MinimumNArgs(1),
	Run:  Register,
}

func init() {
	rootCmd.AddCommand(registerCmd)
	publisherNode = registerCmd.PersistentFlags().BoolP("publisher-node", "p", false, "Register me as a publisher")
}

func Register(cmd *cobra.Command, args []string) {
	nameToregister = args[0]
	fmt.Println("Will register the name ", args[0])
}
