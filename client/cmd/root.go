package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var Password string
var Server string
var registeredName string
var uniqueId string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Client for Device Support",
	Long: `
	Client side support for embedded devices.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.device.yaml)")
	rootCmd.PersistentFlags().StringP("password-salt", "s", "", "Salt for password generation")
	rootCmd.MarkFlagRequired("password-salt")
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
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("../config")
		viper.SetConfigName(".edev")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		Server = viper.GetString("server.url")
		log.Printf("Server URL set to %s", Server)
		registeredName = viper.GetString("client.name")
		uniqueId = viper.GetString("client.uuid")
	} else {
		fmt.Println("No config file. Will use password environment var")
		viper.SetEnvPrefix("edev")
		viper.BindEnv("passsalt")
		viper.BindPFlag("passsalt", rootCmd.PersistentFlags().Lookup("password-salt"))
	}
	Password = viper.GetString("password-salt")
}
