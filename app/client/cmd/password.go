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

var storedLogin string
var storedPassword string

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Adds, shows and changes stored login and password.",
	Long:  `Adds, shows and changes stored login and password.`,
}

var passwordAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new login and password to the storage.",
	Long:  `Add new login and password to the storage.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if recordName == "" {
			log.Fatalln(errors.New("не указано название сохраняемой записи"))
		}

		if storedLogin == "" {
			log.Fatalln(errors.New("не указан логин для сохранения"))
		}

		if storedPassword == "" {
			log.Fatalln(errors.New("не указан пароль для сохранения"))
		}

		for i, arg := range args {
			fmt.Printf("Arg[%d]=%v\n", i, arg)
		}

		keyPhrase := config.GetKeyPhrase()

		encryptor := coder.NewCoder()

		err := encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		encryptedLogin, err := encryptor.Encode([]byte(storedLogin))
		if err != nil {
			log.Fatalln(err)
		}

		encryptedPassword, err := encryptor.Encode([]byte(storedPassword))
		if err != nil {
			log.Fatalln(err)
		}

		addLoginPasswordRequest := srs.AddLoginPasswordRequest{
			Token: &srs.Token{
				Token: token,
			},
			NameMetadata: &srs.RecordNameMetadata{
				Name:     recordName,
				Metadata: recordMetadata,
			},
			EncryptedLoginPassword: &srs.EncryptedLoginPassword{
				EncryptedLogin:    encryptedLogin,
				EncryptedPassword: encryptedPassword,
			},
		}

		addLoginPasswordResponse, err := client.AddLoginPassword(context.Background(), &addLoginPasswordRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Запись сохранена с ID", addLoginPasswordResponse.Id)
	},
}

var passwordShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show stored login and password.",
	Long:  `Show stored login and password with specified ID.`,
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

		getLoginPasswordRequest := srs.GetLoginPasswordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getLoginPasswordResponse, err := client.GetLoginPassword(context.Background(), &getLoginPasswordRequest)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedLogin, err := decryptor.Decode(getLoginPasswordResponse.EncryptedLoginPassword.EncryptedLogin)
		if err != nil {
			log.Fatalln(err)
		}

		decryptedPassword, err := decryptor.Decode(getLoginPasswordResponse.EncryptedLoginPassword.EncryptedPassword)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Полученный логин:", string(decryptedLogin))
		fmt.Println("Полученный пароль:", string(decryptedPassword))
	},
}

var passwordChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change existing login and password.",
	Long:  `Change existing login and password.`,
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

		if storedLogin == "" && storedPassword == "" {
			log.Fatalln(errors.New("не указаны логин и пароль для изменения"))
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Fatalln(err)
		}

		var encryptedLogin, encryptedPassword []byte

		if storedLogin != "" {
			encryptedLogin, err = encryptor.Encode([]byte(storedLogin))
			if err != nil {
				log.Fatalln(err)
			}
		}

		if storedPassword != "" {
			encryptedPassword, err = encryptor.Encode([]byte(storedPassword))
			if err != nil {
				log.Fatalln(err)
			}
		}

		changeLoginPasswordRequest := srs.ChangeLoginPasswordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
			EncryptedLoginPassword: &srs.EncryptedLoginPassword{
				EncryptedLogin:    encryptedLogin,
				EncryptedPassword: encryptedPassword,
			},
		}

		_, err = client.ChangeLoginPassword(context.Background(), &changeLoginPasswordRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Запись с ID %d изменена.\n", id)
	},
}

func init() {
	passwordAddCmd.PersistentFlags().StringVarP(&storedLogin, "login", "l", "", "A login to store")
	passwordAddCmd.PersistentFlags().StringVarP(&storedPassword, "password", "p", "", "A password to store")
	passwordAddCmd.PersistentFlags().StringVarP(&recordName, "name", "n", "", "A name of stored record")
	passwordAddCmd.PersistentFlags().StringVarP(&recordMetadata, "metadata", "m", "", "Metadata of stored record")

	passwordAddCmd.MarkFlagRequired("login")
	passwordAddCmd.MarkFlagRequired("password")
	passwordAddCmd.MarkFlagRequired("name")

	passwordShowCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with login and password")
	passwordShowCmd.MarkFlagRequired("id")

	passwordChangeCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record with login and password")
	passwordChangeCmd.PersistentFlags().StringVarP(&storedLogin, "login", "l", "", "A login to store")
	passwordChangeCmd.PersistentFlags().StringVarP(&storedPassword, "password", "p", "", "A password to store")
	passwordChangeCmd.MarkFlagRequired("id")

	passwordCmd.AddCommand(passwordAddCmd)
	passwordCmd.AddCommand(passwordShowCmd)
	passwordCmd.AddCommand(passwordChangeCmd)
	rootCmd.AddCommand(passwordCmd)
}
