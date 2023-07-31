#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

go test -gcflags "all=-l" -cover -v -race -parallel 2 ./...
