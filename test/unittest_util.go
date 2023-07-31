package test

import (
	"encoding/json"
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
)

func PrepareForHandlerTest() (*echo.Echo, container.Container) {
	conf := createConfig()
	c, _ := container.New(conf)

	e := echo.New()
	e.Validator = c.GetValidator()
	middleware.Init(e, c)
	return e, c
}

func createConfig() *config.Configuration {
	conf := &config.Configuration{}
	conf.Database.Dialect = "sqlite3"
	conf.Database.Host = "file::memory:?cache=shared"
	conf.Database.InitData = true
	conf.Database.Migration = true
	conf.Logs.Encoding = "console"
	conf.Logs.Path = "logs"
	conf.Logs.MaxSize = 50
	conf.Auth.Expires = 72
	return conf
}

// ConvertToString func converts model to string.
func ConvertToString(model interface{}) string {
	bytes, _ := json.Marshal(model)
	return string(bytes)
}

func NewJSONRequest(method, target string, param any) *http.Request {
	req := httptest.NewRequest(method, target, strings.NewReader(ConvertToString(param)))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Add(echo.HeaderAccept, echo.MIMEApplicationJSON)
	return req
}
