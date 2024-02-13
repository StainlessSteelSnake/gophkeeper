// Функции для работы в БД с данными об учётных записях пользователей приложения.

package storage

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	// sqlInsertBankCard содержит SQL-запрос для добавления пользователя приложения.
	sqlInsertUser = `
	INSERT INTO public.users(
	login, password)
	VALUES ($1, $2)`

	// sqlSelectUser содержит SQL-запрос для получения хэша пароля пользователя
	// и количества его записей с зашифрованных данных.
	sqlSelectUser = `
	SELECT u.password, COALESCE(count(r.Id), 0)
	FROM public.users AS u 
	LEFT JOIN public.user_records AS r ON u.login = r.user_login
	WHERE login = $1
	GROUP BY u.password`
)

// AddUser добавляет пользователя приложения.
func (s *Storage) AddUser(ctx context.Context, login, password string) error {
	log.Printf("БД. Добавление пользователя '%s' с хэшем пароля '%v'.\n", login, password)

	if login == "" {
		return errors.New("не указан логин пользователя")
	}

	if password == "" {
		return errors.New("не указан пароль пользователя")
	}

	var pgErr *pgconn.PgError

	ct, err := s.conn.Exec(ctx, sqlInsertUser, login, password)
	if err != nil && !errors.As(err, &pgErr) {
		log.Println("БД. Ошибка при добавлении пользователя:", err)
		return err
	}

	if err != nil && pgErr.Code != pgerrcode.UniqueViolation {
		log.Println("БД. Ошибка при добавлении пользователя, код:", pgErr.Code, ", сообщение:", pgErr.Error())
		return err
	}

	if err != nil {
		log.Println("БД. Ошибка при попытке добавления дублирующего пользователя в БД, код:", pgErr.Code, ", сообщение:", pgErr.Error())
		return err
	}

	log.Println("БД. Добавлено записей пользователей в таблицу:", ct.RowsAffected())

	return nil
}

// GetUser возвращает пароль пользователя приложения и количество его записей о зашифрованных данных.
func (s *Storage) GetUser(ctx context.Context, login string) (passwordHash string, recordCount int, err error) {
	log.Printf("БД. Получение пользователя '%s'.\n", login)

	if login == "" {
		return "", 0, errors.New("не указан логин пользователя")
	}

	row := s.conn.QueryRow(ctx, sqlSelectUser, login)

	err = row.Scan(&passwordHash, &recordCount)
	if err != nil {
		log.Printf("БД. Ошибка при получении пользователя '%s', сообщение: %s.\n", login, err)
		return "", 0, err
	}

	log.Printf("БД. Получение пользователя '%s' завершено успешно.\n", login)
	return passwordHash, recordCount, nil
}
