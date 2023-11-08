package main

import (
	"fmt"
	"log"

	"github.com/StainlessSteelSnake/gophkeeper/app/client/cmd"
	"github.com/StainlessSteelSnake/gophkeeper/app/client/config"
	srs "github.com/StainlessSteelSnake/gophkeeper/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate go env -w GOARCH=386 GOOS=windows
//go:generate go build -o bin/windows/gophkeeper-win32.exe main.go
//go:generate go env -w GOARCH=amd64 GOOS=windows
//go:generate go build -o bin/windows/gophkeeper-win64.exe main.go
//go:generate go env -w GOARCH=386 GOOS=linux
//go:generate go build -o bin/linux/gophkeeper-linux32 main.go
//go:generate go env -w GOARCH=amd64 GOOS=linux
//go:generate go build -o bin/linux/gophkeeper-linux64 main.go
//go:generate go env -w GOARCH=amd64 GOOS=darwin
//go:generate go build -ldflags "-X main.Version=v1.0.1 -X 'main.BuildTime=$(date +'%d.%m.%Y %H:%M:%S')'" -o bin/macos/gophkeeper-mac-i86_64 main.go
//go:generate go env -w GOARCH=arm64 GOOS=darwin
//go:generate go build -o bin/macos/gophkeeper-mac-arm64 main.go

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
