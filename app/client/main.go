package main

//go:generate GOOS=windows GOARCH=386 build -o app/client/bin/windows/gophkeeper-win32 app/client/main.go
//go:generate GOOS=windows GOARCH=amd64 go build -o bin/windows/gophkeeper-win64 app/client/main.go
//go:generate GOOS=linux GOARCH=386 go build -o bin/linux/gophkeeper-linux32 app/client/main.go
//go:generate GOOS=linux GOARCH=amd64 go build -o bin/linux/gophkeeper-linux64 app/client/main.go
//go:generate GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=v1.0.1 -X 'main.BuildTime=$(date +'%d.%m.%Y %H:%M:%S')'" -o app/client/bin/macos/gophkeeper-mac-i86_64 app/client/main.go
//go:generate GOOS=darwin GOARCH=arm64 go build -o bin/macos/gophkeeper-mac-arm64 app/client/main.go

import (
	"fmt"
	"log"

	"github.com/StainlessSteelSnake/gophkeeper/app/client/cmd"
	"github.com/StainlessSteelSnake/gophkeeper/app/client/config"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Version   string
	BuildTime string
)

func main() {
	fmt.Println(Version)
	fmt.Println(BuildTime)

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// установка соединения с gRPC-сервером
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// получение gRPC-интерфейса для установленного соединения
	client := srs.NewGophKeeperClient(conn)

	cmd.Execute(client, cfg)

	/*
		var input []byte = make([]byte, 100)
		for {
			count, err := os.Stdin.Read(input)
			if err == io.EOF {
				log.Println("End of file")
				break
			}

			if err != nil {
				log.Fatalln(err)
				break
			}

			if count == 0 {
				log.Println("0 characters read")
				break
			}

			fmt.Println(string(input))
		}

	*/
}
