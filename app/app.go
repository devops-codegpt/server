package app

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/middleware"
	"github.com/devops-codegpt/server/migration"
	"github.com/devops-codegpt/server/router"
	"github.com/labstack/echo/v4"
)

func App(opts ...CallOption) (*echo.Echo, error) {
	option := newOption(opts...)
	// Load config from config file
	conf, err := config.LoadAppConfig(option.configFile)
	if err != nil {
		return nil, err
	}
	// Get container
	c, err := container.New(conf)
	if err != nil {
		return nil, err
	}

	// Init data
	if err := migration.CreateDatabase(c); err != nil {
		return nil, err
	}
	if err := migration.InitData(c); err != nil {
		return nil, err
	}

	// Setup
	e := echo.New()
	e.Validator = c.GetValidator()
	// Initialize routing
	router.Init(e, c)
	// Whether to allow cross-domain
	if option.cors {
		e.Use(middleware.CorsMiddleware())
	}
	// Initialize middleware
	middleware.Init(e, c)

	return e, nil
}
