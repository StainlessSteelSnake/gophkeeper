package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/StainlessSteelSnake/gophkeeper/internal/storage"
)

// AddBankCard предоставляет сервис сохранения данных банковской карты в зашифрованном виде.
func (s *Server) AddBankCard(ctx context.Context, in *srs.AddBankCardRequest) (*srs.AddBankCardResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса добавления записи о банковской карте (AddBankCard) со входными данными: Token=%s.\n", in.Token.Token)

	var response = srs.AddBankCardResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	bankCard := storage.BankCard{
		CardNumber:  in.EncryptedBankCard.CardNumber,
		CardHolder:  in.EncryptedBankCard.CardHolder,
		ExpiryYear:  in.EncryptedBankCard.ExpiryYear,
		ExpiryMonth: in.EncryptedBankCard.ExpiryMonth,
		Cvc:         in.EncryptedBankCard.Cvc,
	}

	id, err := s.storageController.AddBankCard(
		ctx,
		userLogin,
		in.NameMetadata.Name,
		&bankCard,
		in.NameMetadata.Metadata)

	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка добавления записи о банковской карте: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.Id = int32(id)

	return &response, nil
}

// GetBankCard предоставляет сервис получения сохранённых данных банковской карты в зашифрованном виде.
func (s *Server) GetBankCard(ctx context.Context, in *srs.GetBankCardRequest) (*srs.GetBankCardResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса получения записи о банковской карте (GetBankCard) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var response = srs.GetBankCardResponse{
		EncryptedBankCard: &srs.EncryptedBankCard{},
	}

	bankCard, err := s.storageController.GetBankCard(ctx, userLogin, int(in.Id))
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка получения записи о банковской карте: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	response.EncryptedBankCard.CardNumber = bankCard.CardNumber
	response.EncryptedBankCard.CardHolder = bankCard.CardHolder
	response.EncryptedBankCard.ExpiryYear = bankCard.ExpiryYear
	response.EncryptedBankCard.ExpiryMonth = bankCard.ExpiryMonth
	response.EncryptedBankCard.Cvc = bankCard.Cvc

	return &response, nil
}

// ChangeBankCard предоставляет сервис изменения сохранённых данных банковской карты.
func (s *Server) ChangeBankCard(ctx context.Context, in *srs.ChangeBankCardRequest) (*srs.ChangeBankCardResponse, error) {
	log.Printf("gRPC-Сервер. Вызов сервиса изменения записи о банковской карте (ChangeLoginPassword) со входными данными: Token=%s, ID='%d'.\n", in.Token.Token, in.Id)

	var response = srs.ChangeBankCardResponse{}

	userLogin, _, err := s.authenticator.Authenticate(ctx, in.Token.Token)
	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка авторизации: %s.\n", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var bankCard = new(storage.BankCard)

	if in.EncryptedBankCard.CardNumber == nil ||
		in.EncryptedBankCard.CardHolder == nil ||
		in.EncryptedBankCard.ExpiryYear == nil ||
		in.EncryptedBankCard.ExpiryMonth == nil ||
		in.EncryptedBankCard.Cvc == nil {
		bankCard, err = s.storageController.GetBankCard(ctx, userLogin, int(in.Id))

		if err != nil {
			log.Printf("gRPC-Сервер. Ошибка получения записи о банковской карте: %s.\n", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	if in.EncryptedBankCard.CardNumber != nil {
		bankCard.CardNumber = in.EncryptedBankCard.CardNumber
	}

	if in.EncryptedBankCard.CardHolder != nil {
		bankCard.CardHolder = in.EncryptedBankCard.CardHolder
	}

	if in.EncryptedBankCard.ExpiryYear != nil {
		bankCard.ExpiryYear = in.EncryptedBankCard.ExpiryYear
	}

	if in.EncryptedBankCard.ExpiryMonth != nil {
		bankCard.ExpiryMonth = in.EncryptedBankCard.ExpiryMonth
	}

	if in.EncryptedBankCard.Cvc != nil {
		bankCard.Cvc = in.EncryptedBankCard.Cvc
	}

	err = s.storageController.ChangeBankCard(
		ctx,
		userLogin,
		int(in.Id),
		bankCard)

	if err != nil {
		log.Printf("gRPC-Сервер. Ошибка при изменении записи о банковской карте: %s.\n", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response, nil
}
