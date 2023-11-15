package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	conf "github.com/StainlessSteelSnake/gophkeeper/app/client/config"
	"github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

var ServerAddress string                                                                    // IP-адрес и порт для подключения к серверу приложения.
var config conf.Configurator                                                                // Контроллер параметров приложения.
var clientInit func(cfg conf.Configurator) (services.GophKeeperClient, func() error, error) // Инициализатор подключения к серверу приложения.
var clientClose func() error                                                                // Функция отключения от сервера приложения.
var client services.GophKeeperClient                                                        // gRPC-клиент для передачи данных на сервер приложения.

// rootCmd описывает набор команд клиентского приложения.
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
		var err error
		client, clientClose, err = clientInit(config)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		clientClose()
	},
}

// versionCmd описывает команду для отображения версии, даты и времени сборки клиентского приложения.
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

// init добавляет флаги команд и добавляет сами команды в иерархическую структуру.
func init() {
	rootCmd.PersistentFlags().StringVarP(&ServerAddress, "address", "a", ":3200", "A gRPC server address to connect to.")
	rootCmd.AddCommand(versionCmd)
}

// Execute запускает обработку команд пользователя.
func Execute(initFunc func(cfg conf.Configurator) (services.GophKeeperClient, func() error, error), cfg conf.Configurator) {
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
