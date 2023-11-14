package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/StainlessSteelSnake/gophkeeper/internal/coder"
	"github.com/StainlessSteelSnake/gophkeeper/internal/inout"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

// textCmd описывает набор команд для работы с текстовыми данными.
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Adds, shows and changes stored texts.",
	Long:  `Adds, shows and changes stored texts.`,
}

// textAddCmd описывает команду для сохранения текстовых данных в зашифрованном виде.
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

		text := inout.ReadStringAsBytes(os.Stdin)
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

// textShowCmd описывает команду для получения сохранённых текстовых данных.
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
		inout.WriteBytes(decryptedText, os.Stdout)
	},
}

// textChangeCmd описывает команду для изменения сохранённых текстовых данных.
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

		text := inout.ReadStringAsBytes(os.Stdin)
		if len(text) == 0 {
			log.Fatalln(errors.New("не передан текст для изменения"))
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		encryptedText, err := encryptor.Encode(text)
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

// init добавляет флаги команд и добавляет сами команды в иерархическую структуру.
func init() {
	textAddCmd.PersistentFlags().StringVarP(&recordName, "name", "n", "", "A name of stored record")
	textAddCmd.PersistentFlags().StringVarP(&recordMetadata, "metadata", "m", "", "Metadata of stored record")
	textAddCmd.MarkFlagRequired("name")

	textShowCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with text")
	textShowCmd.MarkFlagRequired("id")

	textChangeCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with text")
	textChangeCmd.MarkFlagRequired("id")

	textCmd.AddCommand(textAddCmd)
	textCmd.AddCommand(textShowCmd)
	textCmd.AddCommand(textChangeCmd)
	rootCmd.AddCommand(textCmd)
}
