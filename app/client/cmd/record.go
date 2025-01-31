package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
)

var recordId string       // Идентификатор записи о сохранённых зашифрованных данных.
var recordName string     // Название записи с сохранёнными зашифрованными данными.
var recordMetadata string // Примечание к записи с сохранёнными зашифрованными данными.

// recordCmd описывает набор команд для работы с сохранёнными записями о зашифрованных данных.
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "CRUD operations with records.",
	Long: `Allows to: 
	- view a list of records, 
	- change name and/or metadata of a record, 
    - delete a record.`,
}

// recordListCmd описывает команду для получения и отображения списка записей о зашифрованных данных.
var recordListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows list of saved records.",
	Long:  `Shows list of saved records.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		token := config.GetToken()
		if token == "" {
			log.Println(errors.New("данные авторизации (токен) не найдены"))
			return
		}

		getUserRecordsRequest := srs.GetUserRecordsRequest{
			Token: &srs.Token{
				Token: token,
			},
		}

		getUserRecordsResponse, err := client.GetUserRecords(context.Background(), &getUserRecordsRequest)
		if err != nil {
			log.Println(err)
			return
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

// recordShowCmd описывает команду для получения и отображения подробной информации о записи.
var recordShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows name, type and metadata of a record.",
	Long:  `Shows name, type and metadata of a record.`,
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

		getUserRecordRequest := srs.GetUserRecordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getUserRecordResponse, err := client.GetUserRecord(context.Background(), &getUserRecordRequest)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("ID:", getUserRecordResponse.UserRecord.Id)
		fmt.Println("Тип хранимых данных:", getUserRecordResponse.UserRecord.RecordType)
		fmt.Println("Название:", getUserRecordResponse.UserRecord.Name)
		fmt.Println("Метаданные:", getUserRecordResponse.UserRecord.Metadata)
	},
}

// recordChangeCmd описывает команду для изменения названия и/или примечания записи.
var recordChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change name and/or metadata of a record.",
	Long:  `Change name and/or metadata of a record.`,
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

		if recordName == "" && recordMetadata == "" {
			log.Println(errors.New("не указаны данные для изменения записи"))
			return
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
			log.Println(err)
			return
		}

		fmt.Printf("Запись с ID '%d' успешно изменена.\n", id)
	},
}

// recordDeleteCmd описывает команду для удаления записи.
var recordDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a record.",
	Long:  `Delete a record.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetToken()
		if token == "" {
			log.Println(errors.New("данные авторизации (токен) не найдены"))
			return
		}

		if recordId == "" {
			log.Println(errors.New("не указан ID удаляемой записи"))
			return
		}

		id, err := strconv.Atoi(recordId)
		if err != nil {
			log.Println("неправильно указан ID удаляемой записи:", err)
			return
		}

		getUserRecordRequest := srs.GetUserRecordRequest{
			Token: &srs.Token{
				Token: token,
			},
			Id: int32(id),
		}

		getUserRecordResponse, err := client.GetUserRecord(context.Background(), &getUserRecordRequest)
		if err != nil {
			log.Println(err)
			return
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
				log.Println(err)
				return
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
			log.Println(err)
			return
		}

		fmt.Printf("Запись с ID '%d' успешно удалена.\n", id)
	},
}

// init добавляет флаги команд и добавляет сами команды в иерархическую структуру.
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
