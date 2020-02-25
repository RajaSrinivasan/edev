package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/RajaSrinivasan/edev/client/impl"
)

var publisherNode *bool

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register the device",
	Long: `
	Register the device with the server to establish 
	a trusted relationship.	
	Argument is the name to register. In most cases, this will be the same as the hostname.`,

	Run: Register,
}

func init() {
	rootCmd.AddCommand(registerCmd)
	publisherNode = registerCmd.PersistentFlags().BoolP("publisher-node", "p", false, "Register me as a publisher")
}

func Register(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		impl.Name = args[0]
		log.Println("Will register the name ", args[0])
	}
	impl.Register(*publisherNode)
}
