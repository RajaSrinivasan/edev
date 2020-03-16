package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/edev/client/impl"
)

var revokeAuthorization bool
var listDevices bool

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Approves, revokes authorization; Lists devices with status",
	Long: `
	When a client (device) registers itself at the server, they enter the "pending" state.
	Till the registration is authorized, requests from the client are not entertained.
	
	The authorization can be revoked by using the revoke option.
	
	The argument is the name of the device to be authorized/revoked.`,
	Run: Authorize,
}

func init() {
	rootCmd.AddCommand(authorizeCmd)
	authorizeCmd.PersistentFlags().BoolVarP(&revokeAuthorization, "revoke", "r", false, "Revoke authorization of client")
	authorizeCmd.PersistentFlags().BoolVarP(&listDevices, "list", "l", false, "Enumerate list of all clients with registration status")
}

func Authorize(cmd *cobra.Command, args []string) {
	fmt.Printf("authorize called Device %s Revoke %v", args[0], revokeAuthorization)
	impl.Authorize(revokeAuthorization, listDevices, args)
}
