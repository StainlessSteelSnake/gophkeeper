package main

import (
	"log"

	"github.com/StainlessSteelSnake/gophkeeper/app/client/cmd"
	"github.com/StainlessSteelSnake/gophkeeper/app/client/config"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Запуск скрипта для сборки исполняемых файлов и передачи в них общих данных:
// - версии приложения;
// - даты и времени сборки.
//go:generate ./generate.sh

var (
	Version   string //
	BuildTime string
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Println(err)
	}

	err = cfg.SetVersion(Version, BuildTime)
	if err != nil {
		log.Fatalln(err)
	}

	cmd.Execute(grpcInit, cfg)
}

func grpcInit(cfg config.Configurator) (srs.GophKeeperClient, func() error) {
	// установка соединения с gRPC-сервером
	conn, err := grpc.Dial(cfg.GetServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	// получение gRPC-интерфейса для установленного соединения
	return srs.NewGophKeeperClient(conn), conn.Close
}
