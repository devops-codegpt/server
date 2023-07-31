package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Custom error codes and error messages
const (
	Ok                  = 200
	NotOk               = 400
	Unauthorized        = 401
	Forbidden           = 403
	InternalServerError = 500
)

const (
	OkMsg                  = "success"
	NotOkMsg               = "failed"
	UnauthorizedMsg        = "login failed. wrong user name or password"
	ForbiddenMsg           = "not authorized to access this resource, please contact the administrator"
	InternalServerErrorMsg = "internal server error"
)

var CustomError = map[int]string{
	Ok:                  OkMsg,
	NotOk:               NotOkMsg,
	Unauthorized:        UnauthorizedMsg,
	Forbidden:           ForbiddenMsg,
	InternalServerError: InternalServerErrorMsg,
}

// ret customizes the unified http response body
type ret struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ret  any    `json:"ret"`
}

func newRet(code int, msg string, data any) *ret {
	return &ret{
		Code: code,
		Msg:  msg,
		Ret:  data,
	}
}

// result send json ret results
func result(c echo.Context, code int, msg string, data any) error {
	resp := newRet(code, msg, data)
	return c.JSON(http.StatusOK, resp)
}

func success(c echo.Context) error {
	return result(c, Ok, OkMsg, echo.Map{})
}

func successWithData(c echo.Context, data any) error {
	return result(c, Ok, OkMsg, data)
}

func failWithMsg(c echo.Context, msg string) error {
	return result(c, NotOk, msg, echo.Map{})
}

func failWithCode(c echo.Context, code int) error {
	msg := NotOkMsg
	if v, ok := CustomError[code]; ok {
		msg = v
	}
	return result(c, code, msg, echo.Map{})
}
