package service

import (
	"context"
	"fmt"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/internal/llm"
	"github.com/devops-codegpt/server/internal/llm/model/replitlm"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/response"
	"github.com/devops-codegpt/server/service/discovery"
)

type CodeService interface {
	RunGenerate(generate *request.CodeGenerate) (*response.CodeGenerate, error)
}

type codeService struct {
	container container.Container
}

const codeLLM = "replitlm"

func (c *codeService) RunGenerate(generate *request.CodeGenerate) (*response.CodeGenerate, error) {
	// Get llm service from service discover center
	discover := discovery.New(c.container)
	replitlmGrpc, err := discover.GetServiceByName(codeLLM)
	if err != nil {
		return nil, err
	}
	replitlmClient := replitlm.New(
		fmt.Sprintf("%s:%d", replitlmGrpc.Address, replitlmGrpc.Port),
		c.container,
	)
	resp, err := replitlmClient.Generate(context.Background(), []string{generate.Prompt}, llm.WithLang(generate.Lang))
	if err != nil {
		return nil, err
	}

	outputCode := make([]string, 0)
	for _, choice := range resp.Choices {
		outputCode = append(outputCode, choice.Text)
	}
	codeResp := &response.CodeGenerate{
		Message: "success",
		Status:  0,
		Result: response.Result{
			Input: response.Input{
				Lang: generate.Lang,
				N:    generate.Num,
				Text: generate.Prompt,
			},
			Output: response.Output{
				Code:               outputCode,
				CompletionTokenNum: resp.Usage.CompletionTokens,
				PromptTokenNum:     resp.Usage.PromptTokens,
			},
		},
	}
	return codeResp, nil
}

func NewCodeService(c container.Container) CodeService {
	return &codeService{
		container: c,
	}
}
