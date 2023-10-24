package main

import (
	"context"
	"log"

	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// установка соединения с gRPC-сервером
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// получение gRPC-интерфейса для установленного соединения
	client := srs.NewGophKeeperClient(conn)

	registerRequest := srs.RegisterRequest{
		LoginPassword: &srs.LoginPassword{
			Login:    "pavlovp83@mail.ru",
			Password: "1qazxsw23edc",
		},
	}

	registerResponse, err := client.Register(context.Background(), &registerRequest)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(registerResponse)

	loginRequest := srs.LoginRequest{
		LoginPassword: &srs.LoginPassword{
			Login:    "pavlovp83@mail.ru",
			Password: "1qazxsw23edc",
		},
	}

	loginResponse, err := client.Login(context.Background(), &loginRequest)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(loginResponse)
}
