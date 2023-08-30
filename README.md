# codegpt-server
[![License](https://img.shields.io/badge/license-Apache_2.0-green)](https://github.com/devops-codegpt/server/blob/main/LICENSE)
[![Tag](https://img.shields.io/badge/tag-v1.0.0-blue)](https://github.com/devops-codegpt/server/tags)




## Introduction

*codegpt-server* is the server of [codegpt](https://github.com/devops-codegpt/) written in Go.



## Prerequisites

- Go >= 1.18.0



## Preparation

### [Consul](https://developer.hashicorp.com/consul/downloads)

- **Install**

```bash
wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
apt update
apt install -y consul
```

- **Run**

```bash
consul agent -dev -ui -client=0.0.0.0
```



### MySQL

- **Deploy**

```bash
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=db_admin mysql:latest
```

- **Init**

```bash
mysql -h 127.0.0.1 -u root -p
mysql> CREATE DATABASE codegpt;
```



## Run
- Make
```bash
make build
```

- Run
```bash
./server
```

> Visit [http://127.0.0.1:8089/api/health](http://127.0.0.1:8089/api/health) in browser to check status as below

```json
{
  "code": 200,
  "ret": "healthy",
  "msg": "success"
}
```



## Usage
- Help
```bash
./server --help
```
output
```
NAME:
   CodeGpt - CodeGpt API for AI ChatCodeOps Service

USAGE:
   server [options]

DESCRIPTION:

   CodeGpt is an AI ChatCodeOps Service.

   The following services are supported:
   - AI Chat
   - AI code analysis
   - AI Devops


COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cors               Whether to allow cross-domain (default: false) [%CORS%]
   --config-file value  Config file [%CONFIG_FILE%]
   --address value      Bind address for the API server. (default: ":8089") [%ADDRESS%]
   --help, -h           show help
```



## Settings

*codegpt-server* parameters can be set in the directory [conf](https://github.com/devops-codegpt/server/blob/main/config/).

An example of configuration in [config.yml](https://github.com/devops-codegpt/server/blob/main/config/config.dev.yml).



## License

Project License can be found [here](LICENSE).



## Reference

- [casbin](https://github.com/casbin/casbin): An authorization library that supports access control models like ACL, RBAC, ABAC in Golang.
- [Consul](https://github.com/hashicorp/consul): a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure.
- [Echo](https://echo.labstack.com/): High performance, extensible, minimalist Go web framework.
- [Gorm](https://github.com/jinzhu/gorm): The fantastic ORM library for Golang.
- [logrus](https://github.com/sirupsen/logrus):  a structured logger for Go (golang), completely API compatible with the standard library logger.
- [lumberjack](https://github.com/natefinch/lumberjack):  a log rolling package for Go.
- [validator](https://github.com/go-playground/validator): Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
- [viper](https://github.com/spf13/viper): Go configuration with fangs.
- [zap](https://github.com/uber-go/zap): Blazing fast, structured, leveled logging in Go.
