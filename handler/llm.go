package handler

import (
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/response"
	"github.com/devops-codegpt/server/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LLMHandler interface {
	GetLLMs(c echo.Context) error
}

type llmHandler struct {
	context container.Container
	service service.LLMService
}

func (l *llmHandler) GetLLMs(c echo.Context) error {
	var req request.LLMList
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}

	llms, err := l.service.FindAllLLMs(&req)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	var resp response.PageData
	resp.PageInfo = req.PageInfo
	resp.List = llms
	return successWithData(c, resp)
}

func NewLLMHandler(c container.Container) LLMHandler {
	return &llmHandler{
		context: c,
		service: service.NewLLMService(c),
	}
}
