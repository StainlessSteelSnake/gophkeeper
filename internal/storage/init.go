package storage

import (
	"context"
	"log"
)

const (
	sqlCreateTableUsers = `
		CREATE TABLE IF NOT EXISTS public.users
		(
		login character varying COLLATE pg_catalog."default" NOT NULL,
		password character varying COLLATE pg_catalog."default" NOT NULL,
		CONSTRAINT users_pkey PRIMARY KEY (login)
		)

		TABLESPACE pg_default;`

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

	sqlCreateTableEncryptedPasswords = `
		CREATE TABLE IF NOT EXISTS public.encrypted_passwords
		(
			Id integer NOT NULL,
			login bytea NOT NULL,
			password bytea NOT NULL,
			CONSTRAINT encrypted_passwords_pkey PRIMARY KEY (Id)
		)

		TABLESPACE pg_default;`

	sqlCreateTableEncryptedBinaries = `
		CREATE TABLE IF NOT EXISTS public.encrypted_binaries
		(
			Id integer NOT NULL,
			binary_data bytea NOT NULL,
			CONSTRAINT encrypted_binaries_pkey PRIMARY KEY (Id)
		)
		
		TABLESPACE pg_default;`

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
