package handler

import (
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CodeHandler interface {
	Generate(c echo.Context) error
}

type codeHandler struct {
	context container.Container
	service service.CodeService
}

func (c2 *codeHandler) Generate(c echo.Context) error {
	var req request.CodeGenerate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	// Params validate
	if err := c.Validate(req); err != nil {
		return err
	}

	// Generate Code
	resp, err := c2.service.RunGenerate(&req)
	if err != nil {
		c2.context.GetLogger().GetZapLogger().Errorf("err: %v", err)
		return c.JSON(http.StatusInternalServerError, req)
	}
	return c.JSON(http.StatusOK, resp)
}

func NewCodeHandler(c container.Container) CodeHandler {
	return &codeHandler{
		context: c,
		service: service.NewCodeService(c),
	}
}
