// Функции для работы в БД с зашифрованными бинарными и текстовыми данными.

package storage

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	// sqlInsertBinary содержит SQL-запрос для добавления записи с зашифрованным текстом или бинарными данными.
	sqlInsertBinary = `
	INSERT INTO public.encrypted_binaries(
	id, binary_data)
	VALUES ($1, $2);
`
	// sqlSelectRecordBinary содержит SQL-запрос для получения записи с зашифрованным текстом или бинарными данными по её идентификатору.
	sqlSelectRecordBinary = `
	SELECT b.binary_data 
	FROM public.encrypted_binaries as b
	INNER JOIN public.user_records as r ON r.id = b.id
	WHERE r.Id = $1 AND r.user_login = $2
`
	// sqlSelectRecordBinary содержит SQL-запрос для изменения записи с зашифрованным текстом или бинарными данными по её идентификатору.
	sqlUpdateRecordBinary = `
	UPDATE public.encrypted_binaries as b
	SET binary_data = $3
	FROM public.user_records as r
	WHERE r.id = b.id AND r.id = $1 AND r.user_login = $2
`
)

// addTextOrBinary добавляет запись о зашифрованных текстовых или бинарных данных пользователя, включая название, тип и метаданные.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого создаётся запись;
// name - название записи;
// binary - последовательность байт, представляющая зашифрованные текстовые или бинарные данные;
// metadata - текстовое примечание к записи;
// recordType - тип добавляемых данных: зашифрованные текстовые данные или зашифрованные бинарные данные.
// Метод возвращает идентификатор добавленной записи.
func (s *Storage) addTextOrBinary(ctx context.Context, userLogin string, name string, binary []byte, metadata string, recordType string) (int, error) {
	if userLogin == "" {
		return 0, errors.New("не указан логин пользователя")
	}

	record := Record{
		UserLogin:  userLogin,
		RecordType: recordType,
		Name:       name,
		Metadata:   metadata,
	}

	id, err := s.addRecord(ctx, &record)
	if err != nil {
		return 0, err
	}

	var pgErr *pgconn.PgError

	_, err = s.conn.Exec(ctx, sqlInsertBinary, id, binary)

	if err != nil && !errors.As(err, &pgErr) {
		log.Printf("БД. Ошибка при добавлении записи в таблицу encrypted_binaries, сообщение: '%s'.\n", err)
		return 0, err
	}

	if err != nil && pgErr.Code != pgerrcode.UniqueViolation {
		log.Printf("БД. Ошибка при добавлении записи в encrypted_binaries encrypted_cards, код '%s', сообщение: '%s'.\n", pgErr.Code, pgErr.Error())
		return 0, err
	}

	if err != nil {
		log.Printf("БД. Ошибка при попытке добавления дублирующей записи в таблицу encrypted_binaries c ID '%d', код '%s', сообщение: '%s'.\n", id, pgErr.Code, pgErr.Error())
		return 0, err
	}

	return id, nil
}

// getTextOrBinary находит и возвращает зашифрованные текстовые или бинарные данные пользователя
// по его имени и идентификатору записи.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого ищется запись;
// id - идентификатор записи.
func (s *Storage) getTextOrBinary(ctx context.Context, userLogin string, id int) ([]byte, error) {
	if userLogin == "" {
		return nil, errors.New("не указан логин пользователя")
	}

	row := s.conn.QueryRow(ctx, sqlSelectRecordBinary, id, userLogin)

	var result []byte

	err := row.Scan(&result)
	if err != nil {
		log.Printf("БД. Ошибка при чтении записи из таблицы encrypted_cards c ID '%d', сообщение: '%s'.\n", id, err)
		return nil, err
	}

	return result, nil
}

// changeTextOrBinary находит и изменяет зашифрованные текстовые или бинарные данные пользователя
// по его имени и идентификатору записи.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого изменяется запись;
// id - идентификатор записи;
// binary- новые зашифрованные текстовые или бинарные данные для изменения.
func (s *Storage) changeTextOrBinary(ctx context.Context, userLogin string, id int, binary []byte) error {
	if userLogin == "" {
		return errors.New("не указан логин пользователя")
	}

	if len(binary) == 0 {
		return errors.New("не переданы данные для изменения")
	}

	_, err := s.conn.Exec(ctx, sqlUpdateRecordBinary, id, userLogin, binary)
	if err != nil {
		log.Printf("БД. Ошибка при попытке обновления записи в таблице encrypted_cards c ID '%d', сообщение: '%s'.\n", id, err)
		return err
	}

	return nil
}

// AddText добавляет запись о зашифрованных текстовых данных пользователя, включая название, тип и метаданные.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого создаётся запись;
// name - название записи;
// text - последовательность байт, представляющая зашифрованные текстовые данные;
// metadata - текстовое примечание к записи.
// Метод возвращает идентификатор добавленной записи.
func (s *Storage) AddText(ctx context.Context, userLogin string, name string, text []byte, metadata string) (int, error) {
	log.Printf("БД. Добавление в таблицу encrypted_binary записи о тексте пользователя '%s' с названием '%v'.\n", userLogin, name)

	id, err := s.addTextOrBinary(ctx, userLogin, name, text, metadata, recordTypeText)
	if err != nil {
		return 0, err
	}

	log.Printf("БД. В таблицу encrypted_binaries добавлена запись о тексте с Id '%d'.\n", id)
	return id, nil
}

// GetText находит и возвращает зашифрованные текстовые данные пользователя
// по его имени и идентификатору записи.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого ищется запись;
// id - идентификатор записи.
func (s *Storage) GetText(ctx context.Context, userLogin string, id int) ([]byte, error) {
	log.Printf("БД. Поиск в таблице encrypted_binaries записи о тексте пользователя '%s' с ID '%d'.\n", userLogin, id)
	return s.getTextOrBinary(ctx, userLogin, id)
}

// ChangeText находит и изменяет зашифрованные текстовые данные пользователя
// по его имени и идентификатору записи.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого изменяется запись;
// id - идентификатор записи;
// text- новые зашифрованные текстовые данные для изменения.
func (s *Storage) ChangeText(ctx context.Context, userLogin string, id int, text []byte) error {
	log.Printf("БД. Обновление записи о тексте в таблице encrypted_binaries, пользователь '%s', ID записи '%d'.\n", userLogin, id)
	return s.changeTextOrBinary(ctx, userLogin, id, text)
}

// AddBinary добавляет запись о зашифрованных бинарных данных пользователя, включая название, тип и метаданные.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого создаётся запись;
// name - название записи;
// binary - последовательность байт, представляющая зашифрованные бинарные данные;
// metadata - текстовое примечание к записи.
// Метод возвращает идентификатор добавленной записи.
func (s *Storage) AddBinary(ctx context.Context, userLogin string, name string, binary []byte, metadata string) (int, error) {
	log.Printf("БД. Добавление в таблицу encrypted_binaries записи о бинарных данных пользователя '%s' с названием '%v'.\n", userLogin, name)

	id, err := s.addTextOrBinary(ctx, userLogin, name, binary, metadata, recordTypeBinary)
	if err != nil {
		return 0, err
	}

	log.Printf("БД. В таблицу encrypted_binaries добавлена запись о бинарных данных с ID '%d'.\n", id)
	return id, nil
}

// GetBinary находит и возвращает зашифрованные бинарные данные пользователя
// по его имени и идентификатору записи.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого ищется запись;
// id - идентификатор записи.
func (s *Storage) GetBinary(ctx context.Context, userLogin string, id int) ([]byte, error) {
	log.Printf("БД. Поиск в таблице encrypted_binaries записи о бинарных данных пользователя '%s' с ID '%d'.\n", userLogin, id)
	return s.getTextOrBinary(ctx, userLogin, id)
}

// ChangeBinary находит и изменяет зашифрованные бинарные данные пользователя
// по его имени и идентификатору записи.
// Входные параметры:
// ctx - контекст для контроля цепочки исполнения программы;
// userLogin - имя пользователя, для которого изменяется запись;
// id - идентификатор записи;
// binary - новые зашифрованные бинарные данные для изменения.
func (s *Storage) ChangeBinary(ctx context.Context, userLogin string, id int, binary []byte) error {
	log.Printf("БД. Обновление записи о бинарных данных в таблице encrypted_binaries, пользователь '%s', ID записи '%d'.\n", userLogin, id)
	return s.changeTextOrBinary(ctx, userLogin, id, binary)
}
