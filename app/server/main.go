package main

import (
	"context"
	"log"

	"github.com/StainlessSteelSnake/gophkeeper/internal/auth"
	"github.com/StainlessSteelSnake/gophkeeper/internal/config"
	"github.com/StainlessSteelSnake/gophkeeper/internal/server"
	"github.com/StainlessSteelSnake/gophkeeper/internal/storage"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfiguration()

	dbStorage := storage.NewStorage(ctx, cfg.DatabaseURI)
	defer dbStorage.Close(ctx)

	authenticator, err := auth.NewAuthentication(dbStorage)
	if err != nil {
		panic(err)
	}

	srv, err := server.NewServer(dbStorage, authenticator, "tcp", cfg.ServerAddress)
	if err != nil {
		panic(err)
	}

	log.Println("gRPC-сервер успешно запущен")

	srv = srv
}
