// Функции для работы в БД с записями о зашифрованных данных пользователя,
// включая названия, типы и текстовые примечания этих записей).

package storage

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	// sqlInsertRecord содержит SQL-запрос для добавления записи о зашифрованных данных пользователя:
	// название, тип данных, примечание.
	sqlInsertRecord = `
	INSERT INTO public.user_records(
	user_login, record_type, name, metadata)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
`
	// sqlSelectUserRecords содержит SQL-запрос для получения списка записей пользователя.
	sqlSelectUserRecords = `
	SELECT id, record_type, name 
	FROM public.user_records
	WHERE user_login = $1
`
	// sqlSelectUserRecordMetadata содержит SQL-запрос для получения примечания к записи пользователя.
	sqlSelectUserRecordMetadata = `
	SELECT id, record_type, name, metadata
	FROM public.user_records
	WHERE id = $1 AND user_login = $2
`
	// sqlUpdateUserRecordName содержит SQL-запрос для изменения названия записи о зашифрованных данных пользователя.
	sqlUpdateUserRecordName = `
	UPDATE public.user_records
	SET name = $3
	WHERE id = $1 AND user_login = $2
`
	// sqlUpdateUserRecordMetadata содержит SQL-запрос для изменения примечания к записи о зашифрованных данных пользователя.
	sqlUpdateUserRecordMetadata = `
	UPDATE public.user_records
	SET metadata = $3
	WHERE id = $1 AND user_login = $2
`
	// sqlUpdateUserRecordNameMetadata содержит SQL-запрос для изменения названия и примечания записи о зашифрованных данных пользователя.
	sqlUpdateUserRecordNameMetadata = `
	UPDATE public.user_records
	SET name = $3, metadata = $4
	WHERE id = $1 AND user_login = $2
`
	// sqlDeleteUserRecord содержит SQL-запрос для удаления записи о зашифрованных данных пользователя.
	sqlDeleteUserRecord = `
	DELETE FROM public.user_records
	WHERE id = $1 AND user_login = $2
`

	// Список типов зашифрованных данных и описаний к этим типам.
	recordTypeLoginPassword            = "LOGIN_PASSWORD" // Логин и пароль
	recordTypeLoginPasswordDescription = "Логин и пароль"
	recordTypeText                     = "TEXT" // Текстовые данные
	recordTypeTextDescription          = "Текст"
	recordTypeBinary                   = "BINARY" // Бинарные данные
	recordTypeBinaryDescription        = "Бинарные данные"
	recordTypeBankCard                 = "BANK_CARD" // Банковская карта
	recordTypeBankCardDescription      = "Банковская карта"
)

// Record хранит запись о зашифрованных данных пользователя.
type Record struct {
	// UserLogin - имя пользователя.
	UserLogin string
	// Id - идентификатор записи.
	Id int
	// RecordType - тип зашифрованных данных.
	RecordType string
	// Name - название записи.
	Name string
	// Metadata - текстовое примечание к записи.
	Metadata string
}

// getRecordType возвращает описание типа зашифрованных данных.
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

// addRecord добавляет запись о зашифрованных данных пользователя.
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

// GetRecords находит и возвращает список записей о зашифрованных данных пользователя по его имени.
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
		if err := rows.Err(); err != nil {
			log.Printf("БД. Ошибка при попытке получения списка записей пользователя '%s', сообщение: '%s'.\n", userLogin, err)
			return nil, err
		}

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

// GetRecord находит и возвращает запись о зашифрованных данных пользователя по его имени и идентификатору записи.
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

// ChangeRecord изменяет запись о зашифрованных данных пользователя.
func (s *Storage) ChangeRecord(ctx context.Context, r *Record) error {
	log.Printf("БД. Изменение записи c ID '%d' для пользователя '%s'.\n", r.Id, r.UserLogin)

	if r.UserLogin == "" {
		return errors.New("не указан логин пользователя")
	}

	if r.Id == 0 {
		return errors.New("не указан идентификатор записи")
	}

	if r.Name == "" && r.Metadata == "" {
		return errors.New("не переданы данные для изменения записи")
	}

	var err error

	if r.Metadata == "" {
		_, err = s.conn.Exec(ctx, sqlUpdateUserRecordName, r.Id, r.UserLogin, r.Name)
	} else if r.Name == "" {
		_, err = s.conn.Exec(ctx, sqlUpdateUserRecordMetadata, r.Id, r.UserLogin, r.Metadata)
	} else {
		_, err = s.conn.Exec(ctx, sqlUpdateUserRecordNameMetadata, r.Id, r.UserLogin, r.Name, r.Metadata)
		fmt.Println(sqlUpdateUserRecordNameMetadata, r.Id, r.UserLogin, r.Name, r.Metadata)
	}

	if err != nil {
		log.Printf("БД. Ошибка при попытке обновления записи в таблице user_records c ID '%d', сообщение: '%s'.\n", r.Id, err)
		return err
	}

	return nil
}

// DeleteRecord изменяет запись о зашифрованных данных пользователя.
func (s *Storage) DeleteRecord(ctx context.Context, userLogin string, id int) error {
	log.Printf("БД. Удаление записи c ID '%d' для пользователя '%s'.\n", id, userLogin)

	if userLogin == "" {
		return errors.New("не указан логин пользователя")
	}

	if id == 0 {
		return errors.New("не указан идентификатор записи")
	}

	_, err := s.conn.Exec(ctx, sqlDeleteUserRecord, id, userLogin)
	if err != nil {
		log.Printf("БД. Ошибка при попытке удаления записи из таблицы user_records c ID '%d', сообщение: '%s'.\n", id, err)
		return err
	}

	return nil
}
