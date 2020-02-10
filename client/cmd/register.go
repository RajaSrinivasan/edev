package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nameToregister, pubkeyfileName, uniqueId *string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register the device",
	Long: `
	Register the device with the server to establish 
	a trusted relationship.

	The unique id is associated with the hostname.
	
	The public key file enables future interactions.
	
	Argument is the name to register. In most cases, this will be the same as the hostname.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Will register the name ", args[0])
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	pubkeyfileName = registerCmd.PersistentFlags().StringP("public-key-filename", "p", "", "Public key filename")
	registerCmd.MarkFlagRequired("public-key-filename")
	uniqueId = registerCmd.PersistentFlags().StringP("unique-id-filename", "u", "", "Unique Id file. If it does not exist, will generate and store")
	registerCmd.MarkFlagRequired("unique-id-filename")

}
