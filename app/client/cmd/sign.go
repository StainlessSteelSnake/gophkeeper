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
		fmt.Println("User:", cmd.Flag("user").Value, ", Password:", cmd.Flag("password").Value, ".")

		registerRequest := srs.RegisterRequest{
			LoginPassword: &srs.LoginPassword{
				Login:    userLogin,    // cmd.Flag("user").Value.String(),
				Password: userPassword, // cmd.Flag("password").Value.String(),
			},
		}

		registerResponse, err := client.Register(context.Background(), &registerRequest)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(registerResponse)

		err = config.SetToken(registerResponse.Token.Token)
		if err != nil {
			log.Fatalln(err)
		}

		keyphrase, err := coder.NewCoder().SetKeyPhrase(userPassword)
		if err != nil {
			log.Fatalln(err)
		}

		err = config.SetKeyPhrase(keyphrase)
		if err != nil {
			log.Fatalln(err)
		}
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
			log.Fatalln(err)
		}
		log.Println(loginResponse)

		err = config.SetToken(loginResponse.Token.Token)
		if err != nil {
			log.Fatalln(err)
		}

		keyphrase, err := coder.NewCoder().SetKeyPhrase(userPassword)
		if err != nil {
			log.Fatalln(err)
		}

		err = config.SetKeyPhrase(keyphrase)
		if err != nil {
			log.Fatalln(err)
		}
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
			log.Fatalln(err)
		}

		err = config.SetKeyPhrase("")
		if err != nil {
			log.Fatalln(err)
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
