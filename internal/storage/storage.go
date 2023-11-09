package storage

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn *pgx.Conn
	user string
}

type Storager interface {
	AddUser(ctx context.Context, login, password string) error
	GetUser(ctx context.Context, login string) (string, int, error)

	GetRecords(ctx context.Context, userLogin string) ([]Record, error)
	GetRecord(ctx context.Context, userLogin string, id int) (*Record, error)
	ChangeRecord(ctx context.Context, r *Record) error
	DeleteRecord(ctx context.Context, userLogin string, id int) error

	AddLoginPassword(ctx context.Context, userLogin string, name string, login []byte, password []byte, metadata string) (int, error)
	GetLoginPassword(ctx context.Context, userLogin string, id int) ([]byte, []byte, error)
	ChangeLoginPassword(ctx context.Context, userLogin string, id int, login []byte, password []byte) error

	AddBankCard(ctx context.Context, userLogin string, name string, bankCard *BankCard, metadata string) (int, error)
	GetBankCard(ctx context.Context, userLogin string, id int) (*BankCard, error)
	ChangeBankCard(ctx context.Context, userLogin string, id int, bankCard *BankCard) error

	AddText(ctx context.Context, userLogin string, name string, text []byte, metadata string) (int, error)
	GetText(ctx context.Context, userLogin string, id int) ([]byte, error)
	ChangeText(ctx context.Context, userLogin string, id int, text []byte) error

	AddBinary(ctx context.Context, userLogin string, name string, binary []byte, metadata string) (int, error)
	GetBinary(ctx context.Context, userLogin string, id int) ([]byte, error)
	ChangeBinary(ctx context.Context, userLogin string, id int, binary []byte) error

	Close(ctx context.Context)
}

func NewStorage(ctx context.Context, databaseUri string) Storager {
	var err error
	storage := &Storage{}

	storage.conn, err = pgx.Connect(ctx, databaseUri)
	if err != nil {
		log.Fatal(err)
		return storage
	}

	dbCfg := strings.Split(databaseUri, ":")
	if len(dbCfg) < 2 {
		log.Fatal(errors.New("в URI базы данных отсутствует информация о пользователе"))
		return storage
	}

	storage.user = strings.TrimPrefix(dbCfg[1], "//")
	if storage.user == "" {
		log.Fatalln(errors.New("в URI базы данных отсутствует информация о пользователе"))
		return storage
	}
	log.Println("Пользователь БД:", storage.user)

	err = storage.init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return storage
}

func (s *Storage) Close(ctx context.Context) {
	if s.conn == nil {
		return
	}

	err := s.conn.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}
}
