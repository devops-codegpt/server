package service

import (
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/models"
	"github.com/devops-codegpt/server/request"
)

type LLMService interface {
	FindAllLLMs(req *request.LLMList) ([]*models.LLM, error)
}

type llmService struct {
	container container.Container
}

func (l *llmService) FindAllLLMs(req *request.LLMList) ([]*models.LLM, error) {
	rep := l.container.GetRepository()
	llms := make([]*models.LLM, 0)
	return models.FindList(rep.Model(&models.LLM{}), &req.PageInfo, llms)
}

func NewLLMService(c container.Container) LLMService {
	return &llmService{container: c}
}
