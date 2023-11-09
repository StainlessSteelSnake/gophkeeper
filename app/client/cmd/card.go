package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/StainlessSteelSnake/gophkeeper/internal/coder"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/spf13/cobra"
)

var storedCardNumber, storedCardHolder, storedCardExpiryYear, storedCardExpiryMonth, storedCardCvc string

var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Adds, shows and changes stored bank card.",
	Long:  `Adds, shows and changes stored bank card.`,
}

var cardAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new bank card to the storage.",
	Long:  `Add new bank card to the storage.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if recordName == "" {
			log.Fatalln(errors.New("не указано название сохраняемой записи"))
		}

		if storedCardNumber == "" {
			log.Fatalln(errors.New("не указан номер банковской карты для сохранения"))
		}

		if storedCardHolder == "" {
			log.Fatalln(errors.New("не указан держатель банковской карты для сохранения"))
		}

		if storedCardExpiryYear == "" {
			log.Fatalln(errors.New("не указан год срока действия банковской карты для сохранения"))
		}

		_, err := checkYear(storedCardExpiryYear)
		if err != nil {
			log.Fatalln(err)
		}

		if storedCardExpiryMonth == "" {
			log.Fatalln(errors.New("не указан месяц срока действия банковской карты для сохранения"))
		}

		_, err = checkMonth(storedCardExpiryMonth)
		if err != nil {
			log.Fatalln(err)
		}

		if storedCardCvc == "" {
			log.Fatalln(errors.New("не указан код CVC банковской карты для сохранения"))
		}

		_, err = checkCvc(storedCardCvc)
		if err != nil {
			log.Fatalln(err)
		}

		for i, arg := range args {
			fmt.Printf("Arg[%d]=%v\n", i, arg)
		}

		keyPhrase := config.GetKeyPhrase()

		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		encryptedCardNumber, err := encryptor.Encode([]byte(storedCardNumber))
		if err != nil {
			log.Fatalln(err)
		}

		encryptedCardHolder, err := encryptor.Encode([]byte(storedCardHolder))
		if err != nil {
			log.Fatalln(err)
		}

		encryptedCardExpiryYear, err := encryptor.Encode([]byte(storedCardExpiryYear))
		if err != nil {
			log.Fatalln(err)
		}

		encryptedCardExpiryMonth, err := encryptor.Encode([]byte(storedCardExpiryMonth))
		if err != nil {
			log.Fatalln(err)
		}

		encryptedCardCvc, err := encryptor.Encode([]byte(storedCardCvc))
		if err != nil {
			log.Fatalln(err)
		}

		addBankCardRequest := srs.AddBankCardRequest{
			Token: &srs.Token{
				Token: token,
			},
			NameMetadata: &srs.RecordNameMetadata{
				Name:     recordName,
				Metadata: recordMetadata,
			},
			EncryptedBankCard: &srs.EncryptedBankCard{
				CardNumber:  encryptedCardNumber,
				CardHolder:  encryptedCardHolder,
				ExpiryYear:  encryptedCardExpiryYear,
				ExpiryMonth: encryptedCardExpiryMonth,
				Cvc:         encryptedCardCvc,
			},
		}

		addBankCardResponse, err := client.AddBankCard(context.Background(), &addBankCardRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Запись сохранена с ID", addBankCardResponse.Id)
	},
}

var cardShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show stored bank card.",
	Long:  `Show stored bank card with specified ID.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if recordId == "" {
			log.Fatalln(errors.New("не указан ID запрашиваемой записи"))
		}

		id, err := strconv.Atoi(recordId)
		if err != nil {
			log.Fatalln("неправильно указан ID запрашиваемой записи:", err)
		}

		keyPhrase := config.GetKeyPhrase()

		decryptor := coder.NewCoder()

		err = decryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		getBankCardRequest := srs.GetBankCardRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getBankCardResponse, err := client.GetBankCard(context.Background(), &getBankCardRequest)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedCardNumber, err := decryptor.Decode(getBankCardResponse.EncryptedBankCard.CardNumber)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedCardHolder, err := decryptor.Decode(getBankCardResponse.EncryptedBankCard.CardHolder)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedCardExpiryYear, err := decryptor.Decode(getBankCardResponse.EncryptedBankCard.ExpiryYear)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedCardExpiryMonth, err := decryptor.Decode(getBankCardResponse.EncryptedBankCard.ExpiryMonth)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedCardCvc, err := decryptor.Decode(getBankCardResponse.EncryptedBankCard.Cvc)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Номер карты:", string(decryptedCardNumber))
		fmt.Println("Держатель карты:", string(decryptedCardHolder))
		fmt.Println("Срок действия:", string(decryptedCardExpiryYear), "/", string(decryptedCardExpiryMonth))
		fmt.Println("CVC/CVV:", string(decryptedCardCvc))
	},
}

var cardChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change existing bank card.",
	Long:  `Change existing bank card.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if recordId == "" {
			log.Fatalln(errors.New("не указан ID изменяемой записи"))
		}

		id, err := strconv.Atoi(recordId)
		if err != nil {
			log.Fatalln("неправильно указан ID запрашиваемой записи:", err)
		}

		if storedCardNumber == "" &&
			storedCardHolder == "" &&
			storedCardExpiryYear == "" &&
			storedCardExpiryMonth == "" &&
			storedCardCvc == "" {
			log.Fatalln(errors.New("не указаны данные банковской карты для изменения"))
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		var encryptedCardNumber, encryptedCardHolder, encryptedCardExpiryYear, encryptedCardExpiryMonth, encryptedCardCvc []byte

		if storedCardNumber != "" {
			encryptedCardNumber, err = encryptor.Encode([]byte(storedCardNumber))
			if err != nil {
				log.Fatalln(err)
			}
		}

		if storedCardHolder != "" {
			encryptedCardHolder, err = encryptor.Encode([]byte(storedCardHolder))
			if err != nil {
				log.Fatalln(err)
			}
		}

		if storedCardExpiryYear != "" {
			_, err = checkYear(storedCardExpiryYear)
			if err != nil {
				log.Fatalln(err)
			}

			encryptedCardExpiryYear, err = encryptor.Encode([]byte(storedCardExpiryYear))
			if err != nil {
				log.Fatalln(err)
			}
		}

		if storedCardExpiryMonth != "" {
			_, err = checkMonth(storedCardExpiryMonth)
			if err != nil {
				log.Fatalln(err)
			}

			encryptedCardExpiryMonth, err = encryptor.Encode([]byte(storedCardExpiryMonth))
			if err != nil {
				log.Fatalln(err)
			}
		}

		if storedCardCvc != "" {
			_, err = checkCvc(storedCardCvc)
			if err != nil {
				log.Fatalln(err)
			}

			encryptedCardCvc, err = encryptor.Encode([]byte(storedCardCvc))
			if err != nil {
				log.Fatalln(err)
			}
		}

		changeBankCardRequest := srs.ChangeBankCardRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
			EncryptedBankCard: &srs.EncryptedBankCard{
				CardNumber:  encryptedCardNumber,
				CardHolder:  encryptedCardHolder,
				ExpiryYear:  encryptedCardExpiryYear,
				ExpiryMonth: encryptedCardExpiryMonth,
				Cvc:         encryptedCardCvc,
			},
		}

		_, err = client.ChangeBankCard(context.Background(), &changeBankCardRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Запись с ID %d изменена.\n", id)
	},
}

func init() {
	cardAddCmd.PersistentFlags().StringVar(&storedCardNumber, "number", "", "A login to store")
	cardAddCmd.PersistentFlags().StringVar(&storedCardHolder, "holder", "", "A password to store")
	cardAddCmd.PersistentFlags().StringVar(&storedCardExpiryYear, "year", "", "A name of stored record")
	cardAddCmd.PersistentFlags().StringVar(&storedCardExpiryMonth, "month", "", "A name of stored record")
	cardAddCmd.PersistentFlags().StringVar(&storedCardCvc, "cvc", "", "A name of stored record")
	cardAddCmd.PersistentFlags().StringVarP(&recordName, "name", "n", "", "A name of stored record")
	cardAddCmd.PersistentFlags().StringVarP(&recordMetadata, "metadata", "m", "", "Metadata of stored record")
	cardAddCmd.MarkFlagRequired("number")
	cardAddCmd.MarkFlagRequired("holder")
	cardAddCmd.MarkFlagRequired("year")
	cardAddCmd.MarkFlagRequired("month")
	cardAddCmd.MarkFlagRequired("cvc")
	cardAddCmd.MarkFlagRequired("name")

	cardShowCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with a bank card data")
	cardShowCmd.MarkFlagRequired("id")

	cardChangeCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with a bank card data")
	cardChangeCmd.PersistentFlags().StringVar(&storedCardNumber, "number", "", "A card number to store")
	cardChangeCmd.PersistentFlags().StringVar(&storedCardHolder, "holder", "", "A card holder to store")
	cardChangeCmd.PersistentFlags().StringVar(&storedCardExpiryYear, "year", "", "An expiration year to store")
	cardChangeCmd.PersistentFlags().StringVar(&storedCardExpiryMonth, "month", "", "An expiration month to store")
	cardChangeCmd.PersistentFlags().StringVar(&storedCardCvc, "cvc", "", "A CVC/CVV to store")
	cardChangeCmd.MarkFlagRequired("id")

	cardCmd.AddCommand(cardAddCmd)
	cardCmd.AddCommand(cardShowCmd)
	cardCmd.AddCommand(cardChangeCmd)
	rootCmd.AddCommand(cardCmd)
}

func checkYear(y string) (int, error) {
	year, err := strconv.Atoi(y)
	if err != nil {
		return 0, errors.New("неправильно указан год срока действия банковской карты")
	}

	if year < 2000 || year > 3000 {
		return 0, errors.New("неправильно указан год срока действия банковской карты")
	}

	return year, nil
}

func checkMonth(m string) (int, error) {
	month, err := strconv.Atoi(m)
	if err != nil {
		return 0, errors.New("неправильно указан месяц срока действия банковской карты")
	}

	if month < 1 || month > 12 {
		return 0, errors.New("неправильно указан месяц срока действия банковской карты")
	}

	return month, nil
}

func checkCvc(c string) (int, error) {
	cvc, err := strconv.Atoi(c)
	if err != nil {
		return 0, errors.New("неправильно указан CVC банковской карты")
	}

	if cvc < 100 || cvc > 999 {
		return 0, errors.New("неправильно указан CVC банковской карты")
	}

	return cvc, nil
}
