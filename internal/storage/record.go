package storage

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	sqlInsertRecord = `
	INSERT INTO public.user_records(
	uuid, user_login, id, record_type, name, metadata)
	VALUES ($1, $2, $3, $4, $5, $6);
`
	sqlInsertLoginPassword = `
	INSERT INTO public.encrypted_passwords(
	uuid, login, password)
	VALUES ($1, $2, $3);
`
	recordTypeLoginPassword = "LOGIN_PASSWORD"
	recordTypeText          = "TEXT"
	recordTypeBinary        = "BINARY"
	recordTypeBankCard      = "BANK_CARD"
)

func getUuid() string {
	uid := uuid.New()
	return hex.EncodeToString(uid[:])
}

func (s *Storage) addRecord(ctx context.Context, user string, id int, recordType string, name string, metadata string) (string, error) {
	log.Printf("БД. Добавление в таблицу user_records записи пользователя '%s' с id '%v', типом '%s' и названием '%v'.\n", user, id, recordType, name)

	var pgErr *pgconn.PgError

	uid := getUuid()
	ct, err := s.conn.Exec(ctx, sqlInsertRecord, uid, user, id, recordType, name, metadata)
	if err != nil && !errors.As(err, &pgErr) {
		log.Println("БД. Ошибка при добавлении записи в таблицу user_records:", err)
		return uid, err
	}

	if err != nil && pgErr.Code != pgerrcode.UniqueViolation {
		log.Println("БД. Ошибка при добавлении записи в таблицу user_records, код:", pgErr.Code, ", сообщение:", pgErr.Error())
		return uid, err
	}

	if err != nil {
		log.Println("БД. Ошибка при попытке добавления дублирующей записи в таблицу user_records, id:", uid, ", код:", pgErr.Code, ", сообщение:", pgErr.Error())
		return uid, err
	}

	log.Println("БД. Добавлено записей с данными в таблицу user_records:", ct.RowsAffected())
	return uid, nil
}

func (s *Storage) AddPassword(ctx context.Context, user string, name string, login []byte, password []byte, metadata string) error {
	_, id, err := s.GetUser(ctx, user)
	if err != nil {
		return err
	}

	id++
	log.Printf("БД. Добавление в таблицу encrypted_passwords записи пользователя '%s' с id '%v', типом 'LOGIN_PASSWORD' и названием '%v'.\n", user, id, name)

	uid, err := s.addRecord(ctx, user, id, recordTypeLoginPassword, name, metadata)
	if err != nil {
		return err
	}

	var pgErr *pgconn.PgError

	ct, err := s.conn.Exec(ctx, sqlInsertLoginPassword, uid, login, password)
	if err != nil && !errors.As(err, &pgErr) {
		log.Println("БД. Ошибка при добавлении записи в таблицу encrypted_passwords:", err)
		return err
	}

	if err != nil && pgErr.Code != pgerrcode.UniqueViolation {
		log.Println("БД. Ошибка при добавлении записи в таблицу encrypted_passwords, код:", pgErr.Code, ", сообщение:", pgErr.Error())
		return err
	}

	if err != nil {
		log.Println("БД. Ошибка при попытке добавления дублирующей записи в таблицу encrypted_passwords, id:", uid, ", код:", pgErr.Code, ", сообщение:", pgErr.Error())
		return err
	}

	log.Println("БД. Добавлено записей с данными в таблицу encrypted_passwords:", ct.RowsAffected())
	return nil
}
