// Функция для создания необходимых для серверного приложения таблиц в БД.

package storage

import (
	"context"
	"log"
)

const (
	// sqlCreateTableUsers - SQL-запрос для создания в БД таблицы с информацией о пользователях приложения.
	sqlCreateTableUsers = `
		CREATE TABLE IF NOT EXISTS public.users
		(
		login character varying COLLATE pg_catalog."default" NOT NULL,
		password character varying COLLATE pg_catalog."default" NOT NULL,
		CONSTRAINT users_pkey PRIMARY KEY (login)
		)

		TABLESPACE pg_default;`

	// sqlCreateTableUserRecords - SQL-запрос для создания в БД таблицы с информацией о записях с зашифрованными данными пользователей.
	// К ним относятся название записи, тип данных и текстовое примечание.
	sqlCreateTableUserRecords = `
		CREATE TABLE IF NOT EXISTS public.user_records
		(
			Id serial NOT NULL,
    		user_login character varying COLLATE pg_catalog."default" NOT NULL,
    		record_type character varying(20) COLLATE pg_catalog."default" NOT NULL,
    		Name character varying COLLATE pg_catalog."default" NOT NULL,
    		Metadata text COLLATE pg_catalog."default",
    		CONSTRAINT user_records_pkey PRIMARY KEY (Id)
		)

		TABLESPACE pg_default;`

	// sqlCreateTableEncryptedPasswords - SQL-запрос для создания в БД таблицы с зашифрованными данными о логинах и паролях.
	sqlCreateTableEncryptedPasswords = `
		CREATE TABLE IF NOT EXISTS public.encrypted_passwords
		(
			Id integer NOT NULL,
			login bytea NOT NULL,
			password bytea NOT NULL,
			CONSTRAINT encrypted_passwords_pkey PRIMARY KEY (Id)
		)

		TABLESPACE pg_default;`

	// sqlCreateTableEncryptedBinaries - SQL-запрос для создания в БД таблицы с зашифрованными текстовыми и бинарными данными.
	sqlCreateTableEncryptedBinaries = `
		CREATE TABLE IF NOT EXISTS public.encrypted_binaries
		(
			Id integer NOT NULL,
			binary_data bytea NOT NULL,
			CONSTRAINT encrypted_binaries_pkey PRIMARY KEY (Id)
		)
		
		TABLESPACE pg_default;`

	// sqlCreateTableEncryptedCards - SQL-запрос для создания в БД таблицы с зашифрованными данными банковских карт.
	sqlCreateTableEncryptedCards = `
		CREATE TABLE IF NOT EXISTS public.encrypted_cards
		(
			Id integer NOT NULL,
			card_number bytea NOT NULL,
			card_holder bytea NOT NULL,
			expiry_year bytea NOT NULL,
			expiry_month bytea NOT NULL,
			cvc bytea NOT NULL,
			CONSTRAINT encrypted_cards_pkey PRIMARY KEY (Id)
		)
		
		TABLESPACE pg_default;`
)

// init последовательно создаёт в БД необходимые для работы приложения таблицы.
func (s *Storage) init(ctx context.Context) error {
	_, err := s.conn.Exec(ctx, sqlCreateTableUsers)
	if err != nil {
		return err
	}

	_, err = s.conn.Exec(ctx, sqlCreateTableUserRecords)
	if err != nil {
		return err
	}

	_, err = s.conn.Exec(ctx, sqlCreateTableEncryptedPasswords)
	if err != nil {
		return err
	}

	_, err = s.conn.Exec(ctx, sqlCreateTableEncryptedBinaries)
	if err != nil {
		return err
	}

	_, err = s.conn.Exec(ctx, sqlCreateTableEncryptedCards)
	if err != nil {
		return err
	}

	log.Println("Таблицы успешно инициализированы в БД")
	return nil
}
