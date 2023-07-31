package replitlm

import (
	"context"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/internal/llm"
	"github.com/devops-codegpt/server/internal/llm/model/replitlm/replitlmclient"
)

type LLM struct {
	client *replitlmclient.Client
}

func (r *LLM) Call(ctx context.Context, prompt string, options ...llm.CallOption) (string, error) {
	resp, err := r.Generate(ctx, []string{prompt}, options...)
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", err
	}
	return resp.Choices[0].Text, nil
}

func (r *LLM) Generate(ctx context.Context, prompts []string, options ...llm.CallOption) (*llm.AIResponse, error) {
	opts := &llm.CallOptions{}
	for _, opt := range options {
		opt(opts)
	}
	resp, err := r.client.RunGenerate(ctx, &replitlmclient.GenerateRequest{
		Prompt: prompts[0],
		Lang:   opts.Lang,
	})
	if err != nil {
		return nil, err
	}
	choices := make([]*llm.Choice, 0)
	for _, code := range resp.CodeList {
		choices = append(choices, &llm.Choice{Text: code})
	}
	return &llm.AIResponse{
		Choices: choices,
		Usage: llm.AIUsage{
			PromptTokens:     resp.PromptTokens,
			CompletionTokens: resp.CompletionTokens,
		},
	}, nil
}

var _ llm.LLM = (*LLM)(nil)

func New(dailTarget string, ctr container.Container) *LLM {
	c := replitlmclient.New(dailTarget, ctr)
	return &LLM{
		client: c,
	}
}
