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

	AddPassword(ctx context.Context, user string, name string, login []byte, password []byte, metadata string) error

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
