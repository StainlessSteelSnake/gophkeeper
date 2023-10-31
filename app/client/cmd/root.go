package cmd

import (
	"os"

	"github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/spf13/cobra"
)

type Configurator interface {
	SetToken(string) error
	GetToken() string
	SetKeyPhrase(string) error
	GetKeyPhrase() string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gophkeeper",
	Short: "GophKeeper is a secure storage for your data",
	Long: `GophKeeper securely store your data in encrypted format on the server. You can use it to keep:
	- a site login and password;
	- a bank card credentials;
	- text, files and so on...`,
}

var client services.GophKeeperClient
var config Configurator

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(cln services.GophKeeperClient, cfg Configurator) {
	if cln == nil || cfg == nil {
		os.Exit(1)
	}
	client = cln
	config = cfg

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
