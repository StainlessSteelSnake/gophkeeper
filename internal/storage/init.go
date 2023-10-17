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
			uuid uuid NOT NULL,
			user_login character varying COLLATE pg_catalog."default" NOT NULL,
			id integer NOT NULL DEFAULT 0,
			record_type character varying(20) COLLATE pg_catalog."default" NOT NULL,
			name character varying COLLATE pg_catalog."default" NOT NULL,
			metadata text COLLATE pg_catalog."default",
			CONSTRAINT user_records_pkey PRIMARY KEY (uuid)
		)

		TABLESPACE pg_default;`

	sqlCreateTableEncryptedPasswords = `
		CREATE TABLE IF NOT EXISTS public.encrypted_passwords
		(
			uuid uuid NOT NULL,
			login bytea NOT NULL,
			password bytea NOT NULL,
			CONSTRAINT encrypted_passwords_pkey PRIMARY KEY (uuid)
		)

		TABLESPACE pg_default;`

	sqlCreateTableEncryptedTexts = `
		CREATE TABLE IF NOT EXISTS public.encrypted_texts
		(
			uuid uuid NOT NULL,
			text_data bytea NOT NULL,
			CONSTRAINT encrypted_texts_pkey PRIMARY KEY (uuid)
		)

		TABLESPACE pg_default;`

	sqlCreateTableEncryptedBinaries = `
		CREATE TABLE IF NOT EXISTS public.encrypted_binaries
		(
			uuid uuid NOT NULL,
			binary_data bytea NOT NULL,
			CONSTRAINT encrypted_binaries_pkey PRIMARY KEY (uuid)
		)
		
		TABLESPACE pg_default;`

	sqlCreateTableEncryptedCards = `
		CREATE TABLE IF NOT EXISTS public.encrypted_cards
		(
			uuid uuid NOT NULL,
			card_number bytea NOT NULL,
			card_holder bytea NOT NULL,
			expiry_year bytea NOT NULL,
			expiry_month bytea NOT NULL,
			cvc bytea NOT NULL,
			CONSTRAINT encrypted_cards_pkey PRIMARY KEY (uuid)
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

	_, err = s.conn.Exec(ctx, sqlCreateTableEncryptedTexts)
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
