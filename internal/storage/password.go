package storage

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	sqlInsertLoginPassword = `
	INSERT INTO public.encrypted_passwords(
	Id, login, password)
	VALUES ($1, $2, $3);
`
	sqlSelectRecordLoginPassword = `
	SELECT lp.login, lp.password 
	FROM public.encrypted_passwords as lp 
	INNER JOIN public.user_records as r ON r.Id = lp.Id
	WHERE r.Id = $1 AND r.user_login = $2
`
)

func (s *Storage) AddLoginPassword(ctx context.Context, userLogin string, name string, login []byte, password []byte, metadata string) (int, error) {
	log.Printf("БД. Добавление в таблицу encrypted_passwords записи пользователя '%s' с названием '%v'.\n", userLogin, name)
	if userLogin == "" {
		return 0, errors.New("не указан логин пользователя")
	}

	record := Record{
		UserLogin:  userLogin,
		RecordType: recordTypeLoginPassword,
		Name:       name,
		Metadata:   metadata,
	}

	id, err := s.addRecord(ctx, &record)
	if err != nil {
		return 0, err
	}

	var pgErr *pgconn.PgError

	_, err = s.conn.Exec(ctx, sqlInsertLoginPassword, id, login, password)

	if err != nil && !errors.As(err, &pgErr) {
		log.Printf("БД. Ошибка при добавлении записи в таблицу encrypted_passwords, сообщение: '%s'.\n", err)
		return 0, err
	}

	if err != nil && pgErr.Code != pgerrcode.UniqueViolation {
		log.Printf("БД. Ошибка при добавлении записи в таблицу encrypted_passwords, код '%s', сообщение: '%s'.\n", pgErr.Code, pgErr.Error())
		return 0, err
	}

	if err != nil {
		log.Printf("БД. Ошибка при попытке добавления дублирующей записи в таблицу encrypted_passwords c Id '%d', код '%s', сообщение: '%s'.\n", id, pgErr.Code, pgErr.Error())
		return 0, err
	}

	log.Printf("БД. В таблицу encrypted_passwords добавлена запись с Id '%d'.\n", id)
	return id, nil
}

func (s *Storage) GetLoginPassword(ctx context.Context, userLogin string, id int) ([]byte, []byte, error) {
	log.Printf("БД. Поиск в таблице encrypted_passwords записи пользователя '%s' с ID '%d'.\n", userLogin, id)
	if userLogin == "" {
		return nil, nil, errors.New("не указан логин пользователя")
	}

	row := s.conn.QueryRow(ctx, sqlSelectRecordLoginPassword, id, userLogin)

	var storedLogin []byte
	var storedPassword []byte

	err := row.Scan(&storedLogin, &storedPassword)
	if err != nil {
		return nil, nil, err
	}

	return storedLogin, storedPassword, nil
}
