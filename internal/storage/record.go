package storage

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	sqlInsertRecord = `
	INSERT INTO public.user_records(
	user_login, record_type, name, Metadata)
	VALUES ($1, $2, $3, $4)
	RETURNING Id;
`
	sqlSelectUserRecords = `
	SELECT Id, record_type, Name 
	FROM public.user_records
	WHERE user_login = $1
`
	sqlSelectUserRecordMetadata = `
	SELECT Id, record_type, Name, Metadata
	FROM public.user_records
	WHERE Id = $1 AND user_login = $2
`

	recordTypeLoginPassword            = "LOGIN_PASSWORD"
	recordTypeLoginPasswordDescription = "Логин и пароль"
	recordTypeText                     = "TEXT"
	recordTypeTextDescription          = "Текст"
	recordTypeBinary                   = "BINARY"
	recordTypeBinaryDescription        = "Бинарные данные"
	recordTypeBankCard                 = "BANK_CARD"
	recordTypeBankCardDescription      = "Банковская карта"
)

type Record struct {
	UserLogin  string
	Id         int
	RecordType string
	Name       string
	Metadata   string
}

func getRecordType(rt string) string {
	switch rt {
	case recordTypeLoginPassword:
		return recordTypeLoginPasswordDescription
	case recordTypeText:
		return recordTypeTextDescription
	case recordTypeBinary:
		return recordTypeBinaryDescription
	case recordTypeBankCard:
		return recordTypeBankCardDescription
	}

	return rt
}

func (s *Storage) addRecord(ctx context.Context, r *Record) (int, error) {
	log.Printf("БД. Добавление в таблицу user_records записи пользователя '%s' с типом '%s' и названием '%v'.\n", r.UserLogin, r.RecordType, r.Name)

	var pgErr *pgconn.PgError

	var recordId int

	row := s.conn.QueryRow(ctx, sqlInsertRecord, r.UserLogin, r.RecordType, r.Name, r.Metadata)
	err := row.Scan(&recordId)

	if err != nil && !errors.As(err, &pgErr) {
		log.Printf("БД. Ошибка при добавлении записи пользователя в таблицу user_records, сообщение: '%s'.\n", err)
		return 0, err
	}

	if err != nil && pgErr.Code != pgerrcode.UniqueViolation {
		log.Printf("БД. Ошибка при добавлении записи пользователя в таблицу user_records, код ошибки '%s', сообщение: '%s'.\n", pgErr.Code, pgErr.Error())
		return 0, err
	}

	if err != nil {
		log.Printf("БД. Ошибка при попытке добавления дублирующей записи в таблицу user_records, код ошибки '%s', сообщение: '%s'.\n", pgErr.Code, pgErr.Error())
		return 0, err
	}

	log.Printf("БД. В таблицу user_records добавленя запись с ID '%d' для пользователя '%s' с типом '%s' и названием '%v'.\n", recordId, r.UserLogin, r.RecordType, r.Name)
	return recordId, nil
}

func (s *Storage) GetRecords(ctx context.Context, userLogin string) ([]Record, error) {
	log.Printf("БД. Получение списка записей пользователя '%s'.\n", userLogin)

	if userLogin == "" {
		return nil, errors.New("не указан логин пользователя")
	}

	result := make([]Record, 0)

	rows, err := s.conn.Query(ctx, sqlSelectUserRecords, userLogin)
	if err == pgx.ErrNoRows {
		log.Printf("Записи для пользователя '%s' не найдены.\n", userLogin)
		return result, nil
	}

	for rows.Next() {
		record := Record{
			UserLogin: userLogin,
		}

		err := rows.Scan(&record.Id, &record.RecordType, &record.Name)
		if err != nil {
			log.Printf("БД. Ошибка при попытке получения списка записей пользователя '%s', сообщение: '%s'.\n", userLogin, err)
			return nil, err
		}

		record.RecordType = getRecordType(record.RecordType)

		result = append(result, record)
	}

	log.Printf("БД. Считан список из %d записей пользователя '%s'.\n", len(result), userLogin)
	return result, nil
}

func (s *Storage) GetRecord(ctx context.Context, userLogin string, id int) (*Record, error) {
	log.Printf("БД. Получение записи c ID '%d' для пользователя '%s'.\n", id, userLogin)

	if userLogin == "" {
		return nil, errors.New("не указан логин пользователя")
	}

	if id == 0 {
		return nil, errors.New("не указан идентификатор записи")
	}

	var result Record

	row := s.conn.QueryRow(ctx, sqlSelectUserRecordMetadata, id, userLogin)

	err := row.Scan(&result.Id, &result.RecordType, &result.Name, &result.Metadata)

	if err != nil {
		log.Printf("БД. Ошибка при попытке получения записи c ID '%d' для пользователя '%s', сообщение: '%s'.\n", id, userLogin, err)
		return nil, err
	}

	result.RecordType = getRecordType(result.RecordType)

	log.Printf("БД. Считана запись с ID '%d' для пользователя '%s'.\n", id, userLogin)
	return &result, nil
}
