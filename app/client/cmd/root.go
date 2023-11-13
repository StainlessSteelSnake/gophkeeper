package cmd

import (
	"fmt"
	"log"
	"os"

	conf "github.com/StainlessSteelSnake/gophkeeper/app/client/config"

	"github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/spf13/cobra"
)

var ServerAddress string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gophkeeper",
	Short: "GophKeeper is a secure storage for your data",
	Long: `GophKeeper securely store your data in encrypted format on the server. You can use it to keep:
	- a site login and password;
	- a bank card credentials;
	- text, files and so on...`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("address").Changed {
			config.SetServerAddress(cmd.Flag("address").Value.String())
		}
		client, clientClose = clientInit(config)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		clientClose()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows current version and build date/time of GophKeeper",
	Long:  "Shows current version and build date/time of GophKeeper.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		version, buildTime := config.GetVersion()

		if version != "" {
			fmt.Println("Версия приложения:", version)
		}

		if buildTime != "" {
			fmt.Println("Дата и время сборки:", buildTime)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
}

var config conf.Configurator
var clientInit func(cfg conf.Configurator) (services.GophKeeperClient, func() error)
var clientClose func() error
var client services.GophKeeperClient

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	rootCmd.PersistentFlags().StringVarP(&ServerAddress, "address", "a", ":3200", "A gRPC server address to connect to.")
	rootCmd.AddCommand(versionCmd)
}

func Execute(initFunc func(cfg conf.Configurator) (services.GophKeeperClient, func() error), cfg conf.Configurator) {
	if initFunc == nil || cfg == nil {
		os.Exit(1)
	}
	config = cfg
	clientInit = initFunc

	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
