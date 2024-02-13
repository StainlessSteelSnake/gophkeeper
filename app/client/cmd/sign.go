package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/StainlessSteelSnake/gophkeeper/internal/coder"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

var userLogin, userPassword string

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Registers, signs in or signs out a user.",
	Long:  `Registers, signs in or signs out a user.`,
}

var signUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Register to GophKeeper server.",
	Long:  `Register to GophKeeper server with user name and password.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		registerRequest := srs.RegisterRequest{
			LoginPassword: &srs.LoginPassword{
				Login:    userLogin,
				Password: userPassword,
			},
		}

		registerResponse, err := client.Register(context.Background(), &registerRequest)
		if err != nil {
			log.Println(err)
			return
		}

		err = config.SetToken(registerResponse.Token.Token)
		if err != nil {
			log.Println(err)
			return
		}

		keyphrase, err := coder.NewCoder().SetKeyPhrase(userPassword)
		if err != nil {
			log.Println(err)
			return
		}

		err = config.SetKeyPhrase(keyphrase)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Пользователь %s успешно зарегистрирован.", userLogin)
	},
}

var signInCmd = &cobra.Command{
	Use:   "in",
	Short: "Login to GophKeeper server.",
	Long:  `Login to GophKeeper server with user name and password.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		loginRequest := srs.LoginRequest{
			LoginPassword: &srs.LoginPassword{
				Login:    userLogin,
				Password: userPassword,
			},
		}

		loginResponse, err := client.Login(context.Background(), &loginRequest)
		if err != nil {
			log.Println(err)
			return
		}

		err = config.SetToken(loginResponse.Token.Token)
		if err != nil {
			log.Println(err)
			return
		}

		keyphrase, err := coder.NewCoder().SetKeyPhrase(userPassword)
		if err != nil {
			log.Println(err)
			return
		}

		err = config.SetKeyPhrase(keyphrase)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Пользователь %s успешно авторизован.", userLogin)
	},
}

var signOut = &cobra.Command{
	Use:   "out",
	Short: "Log out from GophKeeper server.",
	Long:  `Log out from GophKeeper server.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.SetToken("")
		if err != nil {
			log.Println(err)
			return
		}

		err = config.SetKeyPhrase("")
		if err != nil {
			log.Println(err)
			return
		}
	},
}

func init() {
	signCmd.PersistentFlags().StringVarP(&userLogin, "user", "u", "", "The username to log in to the server")
	signCmd.PersistentFlags().StringVarP(&userPassword, "password", "p", "", "The password to log in to the server")

	signCmd.MarkFlagRequired("user")
	signCmd.MarkFlagRequired("password")
	signUpCmd.MarkFlagRequired("user")
	signUpCmd.MarkFlagRequired("password")
	signInCmd.MarkFlagRequired("user")
	signInCmd.MarkFlagRequired("password")

	signCmd.AddCommand(signUpCmd)
	signCmd.AddCommand(signInCmd)
	signCmd.AddCommand(signOut)
	rootCmd.AddCommand(signCmd)
}
