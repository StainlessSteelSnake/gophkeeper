package server

import (
	"context"
	"log"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

	response.UserRecords = make([]*srs.UserRecord, 0)

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
