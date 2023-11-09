package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/StainlessSteelSnake/gophkeeper/internal/coder"
	"github.com/StainlessSteelSnake/gophkeeper/internal/inout"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Adds, shows and changes stored texts.",
	Long:  `Adds, shows and changes stored texts.`,
}

var textAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new text to the storage.",
	Long:  `Add new text to the storage.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if recordName == "" {
			log.Fatalln(errors.New("не указано название сохраняемой записи"))
		}

		text := inout.ReadStringAsBytes()
		if len(text) == 0 {
			log.Fatalln(errors.New("не передан текст для сохранения"))
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err := encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		encryptedText, err := encryptor.Encode(text)
		if err != nil {
			log.Fatalln(err)
		}

		addTextRequest := srs.AddTextRequest{
			Token: &srs.Token{
				Token: token,
			},
			EncryptedText: encryptedText,
			NameMetadata: &srs.RecordNameMetadata{
				Name:     recordName,
				Metadata: recordMetadata,
			},
		}

		addTextResponse, err := client.AddText(context.Background(), &addTextRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Запись сохранена с ID", addTextResponse.Id)
	},
}

var textShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show stored text.",
	Long:  `Show stored text with specified ID.`,
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

		getTextRequest := srs.GetTextRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getTextResponse, err := client.GetText(context.Background(), &getTextRequest)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedText, err := decryptor.Decode(getTextResponse.EncryptedText)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Полученный текст приведён ниже")
		fmt.Println("------------------------------")
		inout.WriteBytes(decryptedText)
	},
}

var textChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change existing text.",
	Long:  `Change existing text.`,
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

		text := inout.ReadStringAsBytes()
		if len(text) == 0 {
			log.Fatalln(errors.New("не передан текст для изменения"))
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		encryptedText, err := encryptor.Encode([]byte(storedLogin))
		if err != nil {
			log.Fatalln(err)
		}

		changeTextRequest := srs.ChangeTextRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id:            int32(id),
			EncryptedText: encryptedText,
		}

		_, err = client.ChangeText(context.Background(), &changeTextRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Запись с ID %d изменена.\n", id)
	},
}

func init() {
	textAddCmd.PersistentFlags().StringVarP(&recordName, "name", "n", "", "A name of stored record")
	textAddCmd.PersistentFlags().StringVarP(&recordMetadata, "metadata", "m", "", "Metadata of stored record")
	textAddCmd.MarkFlagRequired("name")

	textShowCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with login and password")
	textShowCmd.MarkFlagRequired("id")

	textChangeCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with login and password")
	textChangeCmd.MarkFlagRequired("id")

	textCmd.AddCommand(textAddCmd)
	textCmd.AddCommand(textShowCmd)
	textCmd.AddCommand(textChangeCmd)
	rootCmd.AddCommand(textCmd)
}
