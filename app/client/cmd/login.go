package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userLogin, userPassword string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to GophKeeper server. ",
	Long:  `Login to GophKeeper server with user name and password.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("User:", cmd.Flag("user").Value, ", Password:", cmd.Flag("password").Value, ".")
	},
}

func init() {
	loginCmd.Flags().StringVarP(&userLogin, "user", "u", "", "The username to log in to the server")
	loginCmd.Flags().StringVarP(&userPassword, "password", "p", "", "The password to log in to the server")

	rootCmd.AddCommand(loginCmd)
}
