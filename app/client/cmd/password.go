package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/StainlessSteelSnake/gophkeeper/internal/coder"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

var storedLogin string    // Переданный логин для сохранения в зашифрованном виде.
var storedPassword string // Переданный пароль для сохранения в зашифрованном виде.

// passwordCmd описывает набор команд для работы с логинами и паролями.
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Adds, shows and changes stored login and password.",
	Long:  `Adds, shows and changes stored login and password.`,
}

// passwordAddCmd описывает команду для сохранения логина и пароля в зашифрованном виде.
var passwordAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new login and password to the storage.",
	Long:  `Add new login and password to the storage.`,
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

		if storedLogin == "" {
			log.Println(errors.New("не указан логин для сохранения"))
			return
		}

		if storedPassword == "" {
			log.Println(errors.New("не указан пароль для сохранения"))
			return
		}

		for i, arg := range args {
			fmt.Printf("Arg[%d]=%v\n", i, arg)
		}

		keyPhrase := config.GetKeyPhrase()

		encryptor := coder.NewCoder()

		err := encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Println(err)
			return
		}

		encryptedLogin, err := encryptor.Encode([]byte(storedLogin))
		if err != nil {
			log.Println(err)
			return
		}

		encryptedPassword, err := encryptor.Encode([]byte(storedPassword))
		if err != nil {
			log.Println(err)
			return
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
			log.Println(err)
			return
		}

		fmt.Println("Запись сохранена с ID", addLoginPasswordResponse.Id)
	},
}

// passwordShowCmd описывает команду для получения и отображения сохранённых логина и пароля.
var passwordShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show stored login and password.",
	Long:  `Show stored login and password with specified ID.`,
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

		getLoginPasswordRequest := srs.GetLoginPasswordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getLoginPasswordResponse, err := client.GetLoginPassword(context.Background(), &getLoginPasswordRequest)
		if err != nil {
			log.Println(err)
			return
		}

		decryptedLogin, err := decryptor.Decode(getLoginPasswordResponse.EncryptedLoginPassword.EncryptedLogin)
		if err != nil {
			log.Println(err)
			return
		}

		decryptedPassword, err := decryptor.Decode(getLoginPasswordResponse.EncryptedLoginPassword.EncryptedPassword)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Полученный логин:", string(decryptedLogin))
		fmt.Println("Полученный пароль:", string(decryptedPassword))
	},
}

// passwordChangeCmd описывает команду для изменения сохранённых логина и пароля.
var passwordChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change existing login and password.",
	Long:  `Change existing login and password.`,
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

		if storedLogin == "" && storedPassword == "" {
			log.Println(errors.New("не указаны логин и пароль для изменения"))
			return
		}

		keyPhrase := config.GetKeyPhrase()
		encryptor := coder.NewCoder()

		err = encryptor.SetKeyHex(keyPhrase)
		if err != nil {
			log.Println(err)
			return
		}

		var encryptedLogin, encryptedPassword []byte

		if storedLogin != "" {
			encryptedLogin, err = encryptor.Encode([]byte(storedLogin))
			if err != nil {
				log.Println(err)
				return
			}
		}

		if storedPassword != "" {
			encryptedPassword, err = encryptor.Encode([]byte(storedPassword))
			if err != nil {
				log.Println(err)
				return
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
			log.Println(err)
			return
		}

		fmt.Printf("Запись с ID %d изменена.\n", id)
	},
}

// init добавляет флаги команд и добавляет сами команды в иерархическую структуру.
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
