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
			log.Println(errors.New("данные авторизации (токен) не найдены"))
			return
		}

		if recordName == "" {
			log.Println(errors.New("не указано название сохраняемой записи"))
			return
		}

		text := inout.ReadStringAsBytes(os.Stdin)
		if len(text) == 0 {
			log.Println(errors.New("не передан текст для сохранения"))
			return
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err := encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Println(err)
			return
		}

		encryptedText, err := encryptor.Encode(text)
		if err != nil {
			log.Println(err)
			return
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
			log.Println(err)
			return
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
			log.Println(errors.New("данные авторизации (токен) не найдены"))
			return
		}

		if recordId == "" {
			log.Println(errors.New("не указан ID запрашиваемой записи"))
			return
		}

		id, err := strconv.Atoi(recordId)
		if err != nil {
			log.Println("неправильно указан ID запрашиваемой записи:", err)
			return
		}

		keyPhrase := config.GetKeyPhrase()
		decryptor := coder.NewCoder()

		err = decryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Println(err)
			return
		}

		getTextRequest := srs.GetTextRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getTextResponse, err := client.GetText(context.Background(), &getTextRequest)
		if err != nil {
			log.Println(err)
			return
		}

		decryptedText, err := decryptor.Decode(getTextResponse.EncryptedText)
		if err != nil {
			log.Println(err)
			return
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
			log.Println(errors.New("данные авторизации (токен) не найдены"))
			return
		}

		if recordId == "" {
			log.Println(errors.New("не указан ID изменяемой записи"))
			return
		}

		id, err := strconv.Atoi(recordId)
		if err != nil {
			log.Println("неправильно указан ID запрашиваемой записи:", err)
			return
		}

		text := inout.ReadStringAsBytes(os.Stdin)
		if len(text) == 0 {
			log.Println(errors.New("не передан текст для изменения"))
			return
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Println(err)
			return
		}

		encryptedText, err := encryptor.Encode(text)
		if err != nil {
			log.Println(err)
			return
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
			log.Println(err)
			return
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
