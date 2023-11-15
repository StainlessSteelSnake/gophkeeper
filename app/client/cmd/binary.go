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

// binaryCmd описывает набор команд для работы с бинарными данными.
var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Adds, shows and changes stored binary data.",
	Long:  `Adds, shows and changes stored binary data.`,
}

// binaryAddCmd описывает команду для сохранения бинарных данных в зашифрованном виде.
var binaryAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new binary data to the storage.",
	Long:  `Add new binary data to the storage.`,
	Args:  cobra.NoArgs,
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

		binary, err := inout.ReadBytes(os.Stdin)
		if err != nil {
			log.Println(err)
			return
		}

		if len(binary) == 0 {
			log.Println(errors.New("не переданы бинарные данные для сохранения"))
			return
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Println(err)
			return
		}

		encryptedBytes, err := encryptor.Encode(binary)
		if err != nil {
			log.Println(err)
			return
		}

		addBytesRequest := srs.AddBytesRequest{
			Token: &srs.Token{
				Token: token,
			},
			EncryptedBytes: encryptedBytes,
			NameMetadata: &srs.RecordNameMetadata{
				Name:     recordName,
				Metadata: recordMetadata,
			},
		}

		addBytesResponse, err := client.AddBytes(context.Background(), &addBytesRequest)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Запись сохранена с ID", addBytesResponse.Id)
	},
}

// binaryShowCmd описывает команду для получения сохранённых бинарных данных.
var binaryShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show stored binary data.",
	Long:  `Show stored binary data with specified ID.`,
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

		getBytesRequest := srs.GetBytesRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getBytesResponse, err := client.GetBytes(context.Background(), &getBytesRequest)
		if err != nil {
			log.Println(err)
			return
		}

		decryptedBytes, err := decryptor.Decode(getBytesResponse.EncryptedBytes)
		if err != nil {
			log.Println(err)
			return
		}

		inout.WriteBytes(decryptedBytes, os.Stdout)
	},
}

// binaryChangeCmd описывает команду для изменения сохранённых бинарных данных.
var binaryChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change existing binary data.",
	Long:  `Change existing binary data.`,
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

		binary, err := inout.ReadBytes(os.Stdin)
		if err != nil {
			log.Println(err)
			return
		}
		if len(binary) == 0 {
			log.Println(errors.New("не переданы бинарные данные для изменения"))
			return
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Println(err)
			return
		}

		encryptedBytes, err := encryptor.Encode(binary)
		if err != nil {
			log.Println(err)
			return
		}

		changeBytesRequest := srs.ChangeBytesRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id:             int32(id),
			EncryptedBytes: encryptedBytes,
		}

		_, err = client.ChangeBytes(context.Background(), &changeBytesRequest)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Запись с ID %d изменена.\n", id)
	},
}

// init добавляет флаги команд и добавляет сами команды в иерархическую структуру.
func init() {
	binaryAddCmd.PersistentFlags().StringVarP(&recordName, "name", "n", "", "A name of stored record")
	binaryAddCmd.PersistentFlags().StringVarP(&recordMetadata, "metadata", "m", "", "Metadata of stored record")
	binaryAddCmd.MarkFlagRequired("name")

	binaryShowCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with binary data")
	binaryShowCmd.MarkFlagRequired("id")

	binaryChangeCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with binary data")
	binaryChangeCmd.MarkFlagRequired("id")

	binaryCmd.AddCommand(binaryAddCmd)
	binaryCmd.AddCommand(binaryShowCmd)
	binaryCmd.AddCommand(binaryChangeCmd)
	rootCmd.AddCommand(binaryCmd)
}
