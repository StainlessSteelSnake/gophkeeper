package storage

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	sqlInsertBankCard = `
	INSERT INTO public.encrypted_cards(
	id, card_number, card_holder, expiry_year, expiry_month, cvc)
	VALUES ($1, $2, $3, $4, $5, $6);
`
	sqlSelectRecordBankCard = `
	SELECT bc.card_number, bc.card_holder, bc.expiry_year, bc.expiry_month, bc.cvc 
	FROM public.encrypted_cards as bc
	INNER JOIN public.user_records as r ON r.id = bc.id
	WHERE r.id = $1 AND r.user_login = $2
`
	sqlUpdateRecordBankCard = `
	UPDATE public.encrypted_cards as bc
	SET card_number = $3, card_holder = $4, expiry_year = $5, expiry_month = $6, cvc = $7
	FROM public.user_records as r
	WHERE r.id = bc.id AND r.id = $1 AND r.user_login = $2
`
)

type BankCard struct {
	CardNumber  []byte
	CardHolder  []byte
	ExpiryYear  []byte
	ExpiryMonth []byte
	Cvc         []byte
}

func (s *Storage) AddBankCard(ctx context.Context, userLogin string, name string, bankCard *BankCard, metadata string) (int, error) {
	log.Printf("БД. Добавление в таблицу encrypted_cards записи о банковской карте пользователя '%s' с названием '%v'.\n", userLogin, name)
	if userLogin == "" {
		return 0, errors.New("не указан логин пользователя")
	}

	record := Record{
		UserLogin:  userLogin,
		RecordType: recordTypeBankCard,
		Name:       name,
		Metadata:   metadata,
	}

	id, err := s.addRecord(ctx, &record)
	if err != nil {
		return 0, err
	}

	var pgErr *pgconn.PgError

	_, err = s.conn.Exec(
		ctx,
		sqlInsertBankCard,
		id,
		bankCard.CardNumber,
		bankCard.CardHolder,
		bankCard.ExpiryYear,
		bankCard.ExpiryMonth,
		bankCard.Cvc)

	if err != nil && !errors.As(err, &pgErr) {
		log.Printf("БД. Ошибка при добавлении записи о банковской карте в таблицу encrypted_cards, сообщение: '%s'.\n", err)
		return 0, err
	}

	if err != nil && pgErr.Code != pgerrcode.UniqueViolation {
		log.Printf("БД. Ошибка при добавлении записи о банковской карте в таблицу encrypted_cards, код '%s', сообщение: '%s'.\n", pgErr.Code, pgErr.Error())
		return 0, err
	}

	if err != nil {
		log.Printf("БД. Ошибка при попытке добавления дублирующей записи о банковской карте в таблицу encrypted_cards c ID '%d', код '%s', сообщение: '%s'.\n", id, pgErr.Code, pgErr.Error())
		return 0, err
	}

	log.Printf("БД. В таблицу encrypted_cards добавлена запись о банковской карте с ID '%d'.\n", id)
	return id, nil
}

func (s *Storage) GetBankCard(ctx context.Context, userLogin string, id int) (*BankCard, error) {
	log.Printf("БД. Поиск в таблице encrypted_cards записи о банковской карте пользователя '%s' с ID '%d'.\n", userLogin, id)
	if userLogin == "" {
		return nil, errors.New("не указан логин пользователя")
	}

	row := s.conn.QueryRow(ctx, sqlSelectRecordBankCard, id, userLogin)

	var bankCard BankCard

	err := row.Scan(
		&bankCard.CardNumber,
		&bankCard.CardHolder,
		&bankCard.ExpiryYear,
		&bankCard.ExpiryMonth,
		&bankCard.Cvc)
	if err != nil {
		log.Printf("БД. Ошибка при чтении записи о банковской карте из таблицы encrypted_cards c ID '%d', сообщение: '%s'.\n", id, err)
		return nil, err
	}

	return &bankCard, nil
}

func (s *Storage) ChangeBankCard(ctx context.Context, userLogin string, id int, bankCard *BankCard) error {
	log.Printf("БД. Обновление записи о банковской карте в таблице encrypted_cards, пользователь '%s', ID записи '%d'.\n", userLogin, id)
	if userLogin == "" {
		return errors.New("не указан логин пользователя")
	}

	if bankCard == nil {
		return errors.New("не переданы данные о банковской карте")
	}

	_, err := s.conn.Exec(
		ctx,
		sqlUpdateRecordBankCard,
		id,
		userLogin,
		bankCard.CardNumber,
		bankCard.CardHolder,
		bankCard.ExpiryYear,
		bankCard.ExpiryMonth,
		bankCard.Cvc)
	if err != nil {
		log.Printf("БД. Ошибка при попытке обновления записи о банковской карте в таблице encrypted_cards c ID '%d', сообщение: '%s'.\n", id, err)
		return err
	}

	return nil
}
