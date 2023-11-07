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
	log.Printf("gRPC-Сервер. Вызов сервиса регистрации (Register) со входными данными: %v.\n", in.LoginPassword)

	var response = srs.RegisterResponse{
		Token: &srs.Token{},
	}

	token, err := s.authenticator.Register(ctx, in.LoginPassword.Login, in.LoginPassword.Password)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка регистрации: %s.\n", err)
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	response.Token.Token = token
	log.Printf("gRPC-Сервер. Вызов сервиса регистрации (Register) успешен. Получен токен: %s.\n", token)
	return &response, nil
}

func (s *Server) Login(ctx context.Context, in *srs.LoginRequest) (*srs.LoginResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса авторизации (Login) со входными данными: %v.\n", in.LoginPassword)

	var response = srs.LoginResponse{
		Token: &srs.Token{},
	}

	token, err := s.authenticator.Login(ctx, in.LoginPassword.Login, in.LoginPassword.Password)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	response.Token.Token = token
	log.Printf("gRPC-Сервер. Вызов сервиса авторизации (Login) успешен. Получен токен: %s.\n", token)
	return &response, nil
}

func (s *Server) Logout(ctx context.Context, in *srs.LogoutRequest) (*srs.LogoutResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса выхода из учётной записи (Logout) со входными данными: %s.\n", in.Token.Token)

	var response = srs.LogoutResponse{}

	err := s.authenticator.Logout(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &response, nil
}

func (s *Server) GetUserRecords(ctx context.Context, in *srs.GetUserRecordsRequest) (*srs.GetUserRecordsResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса получения списка записей пользователя (GetUserRecords) со входными данными: %s.\n", in.Token.Token)

	var response = srs.GetUserRecordsResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	records, err := s.storageController.GetRecords(ctx, userLogin)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка получения данных о записях пользователя: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.UserRecords = make([]*srs.UserRecord, len(records))

	for _, record := range records {
		userRecord := &srs.UserRecord{
			Id:         int32(record.Id),
			UserLogin:  record.UserLogin,
			RecordType: record.RecordType,
			Name:       record.Name,
			Metadata:   record.Metadata,
		}

		response.UserRecords = append(response.UserRecords, userRecord)
	}

	return &response, nil
}

func (s *Server) GetUserRecord(ctx context.Context, in *srs.GetUserRecordRequest) (*srs.GetUserRecordResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса получения записи пользователя (GetUserRecord) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	var response = srs.GetUserRecordResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	record, err := s.storageController.GetRecord(ctx, userLogin, int(in.Id))
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка получения записи пользователя: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.UserRecord = &srs.UserRecord{
		Id:         int32(record.Id),
		UserLogin:  record.UserLogin,
		RecordType: record.RecordType,
		Name:       record.Name,
		Metadata:   record.Metadata,
	}

	return &response, nil
}

func (s *Server) AddLoginPassword(ctx context.Context, in *srs.AddLoginPasswordRequest) (*srs.AddLoginPasswordResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса добавления записи c логином и паролем (AddLoginPassword) со входными данными: Token=%s.\n", in.Token.Token)

	var response = srs.AddLoginPasswordResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	id, err := s.storageController.AddLoginPassword(
		ctx,
		userLogin,
		in.NameMetadata.Name,
		in.EncryptedLoginPassword.EncryptedLogin,
		in.EncryptedLoginPassword.EncryptedPassword,
		in.NameMetadata.Metadata)

	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка добавления записи c логином и паролем: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.Id = int32(id)

	return &response, nil
}

func (s *Server) GetLoginPassword(ctx context.Context, in *srs.GetLoginPasswordRequest) (*srs.GetLoginPasswordResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса получения записи c логином и паролем (GetLoginPassword) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	var response = srs.GetLoginPasswordResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	encryptedLogin, encryptedPassword, err := s.storageController.GetLoginPassword(ctx, userLogin, int(in.Id))
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка получения записи c логином и паролем: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.EncryptedLogin = encryptedLogin
	response.EncryptedPassword = encryptedPassword

	return &response, nil
}