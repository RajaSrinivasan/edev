package cmd

import (
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/RajaSrinivasan/edev/server/serve"
	device "gitlab.com/RajaSrinivasan/edev/tools"
)

var cfgFile string
var verbosityLevel int
var serverURL string
var deviceDB string

//var serverPort string
var serverCertFileName, pvtKeyFileName string
var htmlPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Embedded device server",
	Long: `
	The server hosted on the cloud that an embedded device relies on
	for communication to the world. The device itself provides no capability
	to be reached unsolicited. `,
	Run: Server,
}

// Server provides the service ie runs as a daemon.
func Server(cmd *cobra.Command, args []string) {
	log.Println("Starting the service")
	serve.ProvideService(serverCertFileName, pvtKeyFileName, serverURL, htmlPath)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")
	rootCmd.PersistentFlags().IntVarP(&verbosityLevel, "verbose", "v", 0, "verbosity level 1 .. 16")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".server" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("../config")
		viper.SetConfigName(".edev")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
		serverURL = viper.GetString("server.url")
		serverPort := viper.GetString("server.port")
		log.Printf("Server URL set to %s", serverURL)

		serverURL = serverURL + ":" + serverPort
		serverCertFileName = viper.GetString("server.certfile")
		pvtKeyFileName = viper.GetString("server.privatekey")
		htmlPath = viper.GetString("server.htmlpath")

		deviceDB = viper.GetString("server.devicedb")
		device.Load(deviceDB)
	}
}
