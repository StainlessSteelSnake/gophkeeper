package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

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
