package main

import (
	"log"

	"github.com/StainlessSteelSnake/gophkeeper/app/client/cmd"
	"github.com/StainlessSteelSnake/gophkeeper/app/client/config"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate ./generate.sh

var (
	Version   string
	BuildTime string
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	cfg.SetVersion(Version, BuildTime)

	// установка соединения с gRPC-сервером
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// получение gRPC-интерфейса для установленного соединения
	client := srs.NewGophKeeperClient(conn)

	cmd.Execute(client, cfg)
}
