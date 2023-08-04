package handler

import (
	"github.com/devops-codegpt/server/container"
	"github.com/labstack/echo/v4"
)

type HealthHandler interface {
	GetHealthCheck(c echo.Context) error
}

type healthHandler struct {
	container container.Container
}

// GetHealthCheck checks whether the api service is active or not
func (h *healthHandler) GetHealthCheck(c echo.Context) error {
	return successWithData(c, "healthy")
}

// NewHealthHandler is constructor
func NewHealthHandler(c container.Container) HealthHandler {
	return &healthHandler{container: c}
}
