package server

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/StainlessSteelSnake/gophkeeper/internal/storage"
)

// AddLoginPassword предоставляет сервис сохранения логина и пароля в зашифрованном виде.
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

// GetLoginPassword предоставляет сервис получения сохранённого логина и пароля в зашифрованном виде.
func (s *Server) GetLoginPassword(ctx context.Context, in *srs.GetLoginPasswordRequest) (*srs.GetLoginPasswordResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса получения записи c логином и паролем (GetLoginPassword) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	var response = srs.GetLoginPasswordResponse{
		EncryptedLoginPassword: &srs.EncryptedLoginPassword{},
	}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	encryptedLogin, encryptedPassword, err := s.storageController.GetLoginPassword(ctx, userLogin, int(in.Id))
	if err != nil && errors.Is(err, storage.ErrorRecordNotFound) {
		log.Printf("gRPC-Сервер. Ошибка получения записи c логином и паролем: %s.\n", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка получения записи c логином и паролем: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.EncryptedLoginPassword.EncryptedLogin = encryptedLogin
	response.EncryptedLoginPassword.EncryptedPassword = encryptedPassword

	return &response, nil
}

// ChangeLoginPassword предоставляет сервис изменения сохранённого логина и пароля.
func (s *Server) ChangeLoginPassword(ctx context.Context, in *srs.ChangeLoginPasswordRequest) (*srs.ChangeLoginPasswordResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса изменения записи c логином и паролем (ChangeLoginPassword) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	var response = srs.ChangeLoginPasswordResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var encryptedLogin, encryptedPassword []byte

	if in.EncryptedLoginPassword.EncryptedLogin == nil || in.EncryptedLoginPassword.EncryptedPassword == nil {
		encryptedLogin, encryptedPassword, err = s.storageController.GetLoginPassword(ctx, userLogin, int(in.Id))
		if err != nil && errors.Is(err, storage.ErrorRecordNotFound) {
			log.Printf("gRPC-Сервер. Ошибка получения записи c логином и паролем: %s.\n", err)
			return nil, status.Error(codes.NotFound, err.Error())
		}

		if err != nil {
			log.Printf("gRPC-Сервер. Ошибка получения записи c логином и паролем: %s.\n", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	if in.EncryptedLoginPassword.EncryptedLogin != nil {
		encryptedLogin = in.EncryptedLoginPassword.EncryptedLogin
	}

	if in.EncryptedLoginPassword.EncryptedPassword != nil {
		encryptedPassword = in.EncryptedLoginPassword.EncryptedPassword
	}

	err = s.storageController.ChangeLoginPassword(
		ctx,
		userLogin,
		int(in.Id),
		encryptedLogin,
		encryptedPassword,
	)

	if err != nil && errors.Is(err, storage.ErrorRecordNotFound) {
		log.Printf("gRPC-Сервер. Ошибка при изменении записи c логином и паролем: %s.\n", err)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка при изменении записи c логином и паролем: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response, nil
}
