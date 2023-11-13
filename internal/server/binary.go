package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

func (s *Server) AddText(ctx context.Context, in *srs.AddTextRequest) (*srs.AddTextResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса добавления записи о тексте (AddText) со входными данными: Token=%s.\n", in.Token.Token)

	var response = srs.AddTextResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	id, err := s.storageController.AddText(
		ctx,
		userLogin,
		in.NameMetadata.Name,
		in.EncryptedText,
		in.NameMetadata.Metadata)

	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка добавления записи о тексте: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.Id = int32(id)

	return &response, nil
}
func (s *Server) GetText(ctx context.Context, in *srs.GetTextRequest) (*srs.GetTextResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса получения записи о тексте (GetText) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var response = srs.GetTextResponse{}

	encryptedText, err := s.storageController.GetText(ctx, userLogin, int(in.Id))
	if err != nil || encryptedText == nil {
		log.Printf("gRPC-Сервер. Ошибка получения записи о тексте: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.EncryptedText = encryptedText

	return &response, nil
}
func (s *Server) ChangeText(ctx context.Context, in *srs.ChangeTextRequest) (*srs.ChangeTextResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса изменения записи о тексте (ChangeText) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var response = srs.ChangeTextResponse{}

	err = s.storageController.ChangeText(ctx, userLogin, int(in.Id), in.EncryptedText)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка при изменении записи о тексте: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response, nil
}
func (s *Server) AddBytes(ctx context.Context, in *srs.AddBytesRequest) (*srs.AddBytesResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса добавления записи о бинарных данных (AddBytes) со входными данными: Token=%s.\n", in.Token.Token)

	var response = srs.AddBytesResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	id, err := s.storageController.AddBinary(
		ctx,
		userLogin,
		in.NameMetadata.Name,
		in.EncryptedBytes,
		in.NameMetadata.Metadata)

	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка добавления записи о бинарных данных: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.Id = int32(id)

	return &response, nil
}
func (s *Server) GetBytes(ctx context.Context, in *srs.GetBytesRequest) (*srs.GetBytesResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса получения записи о бинарных данных (GetBytes) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var response = srs.GetBytesResponse{}

	encryptedBytes, err := s.storageController.GetBinary(ctx, userLogin, int(in.Id))
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка получения записи о бинарных данных: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.EncryptedBytes = encryptedBytes

	return &response, nil
}
func (s *Server) ChangeBytes(ctx context.Context, in *srs.ChangeBytesRequest) (*srs.ChangeBytesResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса изменения записи о бинарных данных (ChangeBytes) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var response = srs.ChangeBytesResponse{}

	err = s.storageController.ChangeText(ctx, userLogin, int(in.Id), in.EncryptedBytes)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка при изменении записи о бинарных данных: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response, nil
}
