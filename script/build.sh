#!/bin/bash

ldflags="-s -w"
os=$(go env GOOS)
target="server"

go env -w GOPROXY=https://goproxy.cn,direct

if [ "$os" = "windows" ]; then
  GOARCH=amd64 GOOS=windows go build -ldflags "$ldflags" -o bin/$target.exe main.go
  upx bin/$target.exe
else
  GOARCH=amd64 GOOS=linux go build -ldflags "$ldflags" -o bin/$target main.go
  upx bin/$target
fi
