#!/bin/zsh

version='v1.0.1'
build_time=$(date +'%Y/%m/% %H:%M:%S')

go env -w GOARCH=amd64 GOOS=darwin
go build -ldflags "-X main.Version=$version -X 'main.BuildTime=$build_time'" -o bin/macos/gophkeeper-mac-i86_64 main.go
go env -w GOARCH=arm64 GOOS=darwin
go build -ldflags "-X main.Version=$version -X 'main.BuildTime=$build_time'" -o bin/macos/gophkeeper-mac-arm64 main.go
go env -w GOARCH=386 GOOS=windows
go build -ldflags "-X main.Version=$version -X 'main.BuildTime=$build_time'" -o bin/windows/gophkeeper-win32.exe main.go
go env -w GOARCH=amd64 GOOS=windows
go build -ldflags "-X main.Version=$version -X 'main.BuildTime=$build_time'" -o bin/windows/gophkeeper-win64.exe main.go
go env -w GOARCH=386 GOOS=linux
go build -ldflags "-X main.Version=$version -X 'main.BuildTime=$build_time'" -o bin/linux/gophkeeper-linux32 main.go
go env -w GOARCH=amd64 GOOS=linux
go build -ldflags "-X main.Version=$version -X 'main.BuildTime=$build_time'" -o bin/linux/gophkeeper-linux64 main.go
go env -u GOARCH GOOS

echo "Binaries generation is completed."