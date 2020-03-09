package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var revokeAuthorization *bool

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Authorizes a pending registration",
	Long: `
	When a client (device) registers itself at the server, they enter the "pending" state.
	Till the registration is authorize, requests from the client are not entertained.
	
	The authorization can be revoked by using the revoke option.
	
	The argument is the name of the device to be authorized/revoked.`,
	Run:  Authorize,
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(authorizeCmd)
	authorizeCmd.PersistentFlags().BoolVarP(&revokeAuthorization, "revoke", "r", false, "Revoke authorization of client")
}

func Authorize(cmd *cobra.Command, args []string) {
	fmt.Println("authorize called Device %s Revoke %v", args[0], revokeAuthorization)

}
