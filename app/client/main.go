package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/StainlessSteelSnake/gophkeeper/app/client/cmd"
	"github.com/StainlessSteelSnake/gophkeeper/app/client/config"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

// Запуск скрипта для сборки исполняемых файлов и передачи в них общих данных:
// - версии приложения;
// - даты и времени сборки.
//go:generate ./generate.sh

var (
	Version   string // Версия клиентского приложения.
	BuildTime string // Дата и время сборки клиентского приложения.
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Println(err)
	}

	err = cfg.SetVersion(Version, BuildTime)
	if err != nil {
		log.Println(err)
		return
	}

	cmd.Execute(grpcInit, cfg)
}

// grpcInit создаёт gRPC-клиент и устанавливает соединение с сервером приложения.
func grpcInit(cfg config.Configurator) (srs.GophKeeperClient, func() error, error) {
	// установка соединения с gRPC-сервером
	conn, err := grpc.Dial(cfg.GetServerAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	// получение gRPC-интерфейса для установленного соединения
	return srs.NewGophKeeperClient(conn), conn.Close, nil
}
