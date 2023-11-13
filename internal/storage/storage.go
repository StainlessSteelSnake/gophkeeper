// Функции для получения экземпляра контроллера БД.
// Контроллер представляет интерфейс серверу приложения для работы с сохраняемыми данными пользователей.

package storage

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
)

// Storage представляет контроллер для работы с БД на уровне сервера приложения.
// Хранит информацию о соединении с БД и о пользователе БД.
type Storage struct {
	conn *pgx.Conn
	user string
}

// Storager представляет функции контроллера БД для добавления, получения, изменения и удаления данных.
type Storager interface {
	// AddUser добавляет пользователя приложения.
	AddUser(ctx context.Context, login, password string) error
	// GetUser возвращает пароль пользователя приложения и количество его записей о зашифрованных данных.
	GetUser(ctx context.Context, login string) (string, int, error)

	// GetRecords возвращает список записей о зашифрованных данных пользователя.
	GetRecords(ctx context.Context, userLogin string) ([]Record, error)
	// GetRecord возвращает запись о зашифрованных данных пользователя по её индентификатору.
	GetRecord(ctx context.Context, userLogin string, id int) (*Record, error)
	// ChangeRecord изменяет запись о зашифрованных данных пользователя.
	ChangeRecord(ctx context.Context, r *Record) error
	// DeleteRecord удаляет запись о зашифрованных данных пользователя.
	DeleteRecord(ctx context.Context, userLogin string, id int) error

	// AddLoginPassword добавляет запись о зашифрованных данных о логине и пароле пользователя.
	AddLoginPassword(ctx context.Context, userLogin string, name string, login []byte, password []byte, metadata string) (int, error)
	// GetLoginPassword находит и возвращает зашифрованные данные о логине и пароле пользователя.
	GetLoginPassword(ctx context.Context, userLogin string, id int) ([]byte, []byte, error)
	// ChangeLoginPassword находит и изменяет зашифрованные данные о логине и пароле пользователя.
	ChangeLoginPassword(ctx context.Context, userLogin string, id int, login []byte, password []byte) error

	// AddBankCard добавляет запись о зашифрованных данных банковской карты пользователя.
	AddBankCard(ctx context.Context, userLogin string, name string, bankCard *BankCard, metadata string) (int, error)
	// GetBankCard находит и возвращает зашифрованные данные о банковской карте пользователя.
	GetBankCard(ctx context.Context, userLogin string, id int) (*BankCard, error)
	// ChangeBankCard находит и изменяет зашифрованные данные о банковской карте пользователя.
	ChangeBankCard(ctx context.Context, userLogin string, id int, bankCard *BankCard) error

	// AddText добавляет запись о зашифрованных текстовых данных пользователя.
	AddText(ctx context.Context, userLogin string, name string, text []byte, metadata string) (int, error)
	// GetText находит и возвращает зашифрованные текстовые данные пользователя.
	GetText(ctx context.Context, userLogin string, id int) ([]byte, error)
	// ChangeText находит и изменяет зашифрованные текстовые данные пользователя.
	ChangeText(ctx context.Context, userLogin string, id int, text []byte) error

	// AddBinary добавляет запись о зашифрованных бинарных данных пользователя.
	AddBinary(ctx context.Context, userLogin string, name string, binary []byte, metadata string) (int, error)
	// GetBinary находит и возвращает зашифрованные бинарные данные пользователя.
	GetBinary(ctx context.Context, userLogin string, id int) ([]byte, error)
	// ChangeBinary находит и изменяет зашифрованные бинарные данные пользователя.
	ChangeBinary(ctx context.Context, userLogin string, id int, binary []byte) error

	// Close закрывает соединение с БД.
	Close(ctx context.Context)
}

// ErrorRecordNotFound содержит универсальную ошибку для случаев,
// когда искомая запись о зашифрованных данных не найдена.
var ErrorRecordNotFound = errors.New("Запись с указанным идентификатором не найдена")

// NewStorage создаёт контроллер хранилища, устанавливает соединение с БД.
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

// Close закрывает соединение с БД.
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
