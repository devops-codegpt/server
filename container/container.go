// Package container accesses the data which sharing in overall application.
package container

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/container/authentication"
	"github.com/devops-codegpt/server/container/logger"
	"github.com/devops-codegpt/server/container/repository"
	"github.com/devops-codegpt/server/container/validator"
)

// Container represents an interface for accessing the data which sharing in overall application.
type Container interface {
	GetConfig() *config.Configuration
	GetRepository() repository.Repository
	GetLogger() logger.Logger
	GetAuthentication() authentication.Authentication
	GetValidator() validator.CustomValidator
}

// container struct is for sharing data which such as database setting, the setting of application and logger in overall this application.
type container struct {
	config    *config.Configuration
	logger    logger.Logger
	rep       repository.Repository
	auth      authentication.Authentication
	validator validator.CustomValidator
}

// GetRepository returns the object of repository
func (c *container) GetRepository() repository.Repository {
	return c.rep
}

// GetLogger returns the object of logger
func (c *container) GetLogger() logger.Logger {
	return c.logger
}

// GetConfig returns the object of configuration.
func (c *container) GetConfig() *config.Configuration {
	return c.config
}

func (c *container) GetAuthentication() authentication.Authentication {
	return c.auth
}

func (c *container) GetValidator() validator.CustomValidator {
	return c.validator
}

// New is constructor.
func New(conf *config.Configuration) (Container, error) {
	// Init log
	log, err := logger.InitLogger(conf)
	if err != nil {
		return nil, err
	}

	// Init repository
	rep, err := repository.NewRepository(log, conf)
	if err != nil {
		return nil, err
	}

	// Load authentication rules
	auth, err := authentication.NewAuthentication(rep)
	if err != nil {
		return nil, err
	}

	// Load global validator
	v := validator.New()

	return &container{
		config:    conf,
		rep:       rep,
		logger:    log,
		auth:      auth,
		validator: v,
	}, nil
}
