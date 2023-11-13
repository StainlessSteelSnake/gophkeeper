package server

import (
	"errors"
	"net"

	"google.golang.org/grpc"

	"github.com/StainlessSteelSnake/gophkeeper/internal/auth"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/StainlessSteelSnake/gophkeeper/internal/storage"
)

type Server struct {
	srs.UnimplementedGophKeeperServer
	storageController storage.Storager
	authenticator     auth.Authenticator
}

func NewServer(storageController storage.Storager, authenticator auth.Authenticator, network, address string) (*Server, error) {
	if storageController == nil {
		return nil, errors.New("не задан контроллер хранилища данных")
	}

	if authenticator == nil {
		return nil, errors.New("не задан контроллер регистрации и идентификации пользователей")
	}

	server := Server{
		storageController: storageController,
		authenticator:     authenticator,
	}

	// определяем порт для сервера
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	// создаём gRPC-сервер без зарегистрированной службы
	grpcServer := grpc.NewServer()
	// регистрируем сервис
	srs.RegisterGophKeeperServer(grpcServer, &server)

	// обрабатываем gRPC-запросы
	err = grpcServer.Serve(listener)
	if err != nil {
		return nil, err
	}

	return &server, nil
}
