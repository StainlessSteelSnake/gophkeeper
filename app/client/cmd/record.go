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

var recordId string
var recordName string
var recordMetadata string

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
	Short: "Shows name, type and metadata of a record.",
	Long:  `Shows name, type and metadata of a record.`,
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

		getUserRecordRequest := srs.GetUserRecordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
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

var recordChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change name and/or metadata of a record.",
	Long:  `Change name and/or metadata of a record.`,
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

		if recordName == "" && recordMetadata == "" {
			log.Fatalln(errors.New("не указаны данные для изменения записи"))
		}

		changeUserRecordRequest := srs.ChangeUserRecordRequest{
			Token: &srs.Token{
				Token: token,
			},
			UserRecord: &srs.UserRecord{
				Id:       int32(id),
				Name:     recordName,
				Metadata: recordMetadata,
			},
		}

		_, err = client.ChangeUserRecord(context.Background(), &changeUserRecordRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Запись с ID '%d' успешно изменена.\n", id)
	},
}

var recordDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a record.",
	Long:  `Delete a record.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Fatalln(errors.New("данные авторизации (токен) не найдены"))
		}

		if recordId == "" {
			log.Fatalln(errors.New("не указан ID удаляемой записи"))
		}

		id, err := strconv.Atoi(recordId)
		if err != nil {
			log.Fatalln("неправильно указан ID удаляемой записи:", err)
		}

		getUserRecordRequest := srs.GetUserRecordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getUserRecordResponse, err := client.GetUserRecord(context.Background(), &getUserRecordRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf(
			"Подтвердить удаление записи с ID '%d' название '%s' и типом данных '%s'? (Y/N) (Д/Н)\n",
			getUserRecordResponse.UserRecord.Id,
			getUserRecordResponse.UserRecord.Name,
			getUserRecordResponse.UserRecord.RecordType)

		var confirmed bool
		for !confirmed {
			var answer string
			_, err := fmt.Scan(&answer)
			if err != nil {
				log.Fatalln(err)
			}

			switch answer {
			case "n", "N", "Н", "н":
				fmt.Printf("Удаление записи с ID '%d' не подтверждено.\n", getUserRecordResponse.UserRecord.Id)
				return
			case "y", "Y", "Д", "д":
				fmt.Printf("Удаление записи с ID '%d' подтверждено.\n", getUserRecordResponse.UserRecord.Id)
				confirmed = true
			}
		}

		deleteUserRecordRequest := srs.DeleteUserRecordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		_, err = client.DeleteUserRecord(context.Background(), &deleteUserRecordRequest)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Запись с ID '%d' успешно удалена.\n", id)
	},
}

func init() {
	recordShowCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record")
	recordShowCmd.MarkFlagRequired("id")

	recordChangeCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record")
	recordChangeCmd.PersistentFlags().StringVarP(&recordName, "name", "n", "", "A name of stored record")
	recordChangeCmd.PersistentFlags().StringVarP(&recordMetadata, "metadata", "m", "", "Metadata of stored record")
	recordChangeCmd.MarkFlagRequired("id")

	recordDeleteCmd.PersistentFlags().StringVarP(&recordId, "id", "i", "", "The ID of required record")
	recordDeleteCmd.MarkFlagRequired("id")

	recordCmd.AddCommand(recordListCmd)
	recordCmd.AddCommand(recordShowCmd)
	recordCmd.AddCommand(recordChangeCmd)
	recordCmd.AddCommand(recordDeleteCmd)
	rootCmd.AddCommand(recordCmd)
}
