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

var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Adds, shows and changes stored binary data.",
	Long:  `Adds, shows and changes stored binary data.`,
}

var binaryAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new binary data to the storage.",
	Long:  `Add new binary data to the storage.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if recordName == "" {
			log.Fatalln(errors.New("не указано название сохраняемой записи"))
		}

		binary, err := inout.ReadBytes()
		if err != nil {
			log.Fatalln(err)
		}

		if len(binary) == 0 {
			log.Fatalln(errors.New("не переданы бинарные данные для сохранения"))
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		encryptedBytes, err := encryptor.Encode(binary)
		if err != nil {
			log.Fatalln(err)
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
			log.Fatalln(err)
		}

		fmt.Println("Запись сохранена с ID", addBytesResponse.Id)
	},
}

var binaryShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show stored binary data.",
	Long:  `Show stored binary data with specified ID.`,
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

		getBytesRequest := srs.GetBytesRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getBytesResponse, err := client.GetBytes(context.Background(), &getBytesRequest)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedBytes, err := decryptor.Decode(getBytesResponse.EncryptedBytes)
		if err != nil {
			log.Fatalln(err)
		}

		inout.WriteBytes(decryptedBytes)
	},
}

var binaryChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change existing binary data.",
	Long:  `Change existing binary data.`,
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

		binary, err := inout.ReadBytes()
		if err != nil {
			log.Fatalln(err)
		}
		if len(binary) == 0 {
			log.Fatalln(errors.New("не переданы бинарные данные для изменения"))
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		encryptedBytes, err := encryptor.Encode(binary)
		if err != nil {
			log.Fatalln(err)
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
			log.Fatalln(err)
		}

		fmt.Printf("Запись с ID %d изменена.\n", id)
	},
}

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
