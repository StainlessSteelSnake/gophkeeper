package server

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/StainlessSteelSnake/gophkeeper/internal/auth"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/StainlessSteelSnake/gophkeeper/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) Register(ctx context.Context, in *srs.RegisterRequest) (*srs.RegisterResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса регистрации (Register) со входными данными: %v, \n", in.LoginPassword)

	var response = srs.RegisterResponse{
		Token: &srs.Token{},
	}

	token, err := s.authenticator.Register(ctx, in.LoginPassword.Login, in.LoginPassword.Password)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка регистрации: %v, \n", err)
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	response.Token.Token = token
	log.Printf("gRPC-Сервер. Вызов сервиса регистрации (Register) успешен. Получен токен: %v, \n", token)
	return &response, nil
}

func (s *Server) Login(ctx context.Context, in *srs.LoginRequest) (*srs.LoginResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса авторизации (Login) со входными данными: %v, \n", in.LoginPassword)

	var response = srs.LoginResponse{
		Token: &srs.Token{},
	}

	token, err := s.authenticator.Login(ctx, in.LoginPassword.Login, in.LoginPassword.Password)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %v, \n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	response.Token.Token = token
	log.Printf("gRPC-Сервер. Вызов сервиса авторизации (Login) успешен. Получен токен: %v, \n", token)
	return &response, nil
}
