package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"github.com/spf13/cobra"
)

var id string

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "CRUD operations with records.",
	Long: `Allows to: 
	- view a list of records, 
	- change name and/or metadata of a record, 
    - delete a record.`,
}

var recordListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows list of saved records.",
	Long:  `Shows list of saved records.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		getUserRecordsRequest := srs.GetUserRecordsRequest{
			Token: &srs.Token{
				Token: token,
			},
		}

		getUserRecordsResponse, err := client.GetUserRecords(context.Background(), &getUserRecordsRequest)
		if err != nil {
			log.Fatalln(err)
		}

		if len(getUserRecordsResponse.UserRecords) == 0 {
			fmt.Println("Сохранённые данные не найдены.")
		} else {
			fmt.Printf("ID | Тип хранимых данных | Название\n")
		}

		for _, userRecord := range getUserRecordsResponse.UserRecords {
			fmt.Printf("%d | %s | %s\n", userRecord.Id, userRecord.RecordType, userRecord.Name)
		}

	},
}

var recordShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows list of saved records.",
	Long:  `Shows list of saved records.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if id == "" {
			log.Fatalln(errors.New("не указан ID запрашиваемой записи"))
		}

		recordId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln("неправильно указан ID запрашиваемой записи:", err)
		}

		getUserRecordRequest := srs.GetUserRecordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(recordId),
		}

		getUserRecordResponse, err := client.GetUserRecord(context.Background(), &getUserRecordRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("ID:", getUserRecordResponse.UserRecord.Id)
		fmt.Println("Тип хранимых данных:", getUserRecordResponse.UserRecord.RecordType)
		fmt.Println("Название:", getUserRecordResponse.UserRecord.Name)
		fmt.Println("Метаданные:", getUserRecordResponse.UserRecord.Metadata)
	},
}

func init() {
	recordShowCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "The password to log in to the server")

	recordShowCmd.MarkFlagRequired("id")

	recordCmd.AddCommand(recordListCmd)
	recordCmd.AddCommand(recordShowCmd)
	rootCmd.AddCommand(recordCmd)
}
