package openai

import (
	"context"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/internal/llm"
	"github.com/devops-codegpt/server/internal/llm/model/openai/openaiclient"
)

type LLM struct {
	client   *openaiclient.Client
	callback openaiclient.Callback
	stream   bool
}

func (l *LLM) Call(ctx context.Context, prompt string, options ...llm.CallOption) (string, error) {
	r, err := l.Generate(ctx, []string{prompt}, options...)
	if err != nil {
		return "", err
	}
	if len(r.Choices) == 0 {
		return "", err
	}
	return r.Choices[0].Text, nil
}

func (l *LLM) Generate(ctx context.Context, prompts []string, options ...llm.CallOption) (*llm.AIResponse, error) {
	opts := &llm.CallOptions{}
	for _, opt := range options {
		opt(opts)
	}
	resp, err := l.client.CreateChatWithStream(ctx, &openaiclient.ChatRequest{
		Model:       opts.Model,
		Prompt:      prompts[0],
		Temperature: opts.Temperature,
	})
	if err != nil {
		return nil, err
	}
	return &llm.AIResponse{
		Choices: []*llm.Choice{
			{Text: resp},
		},
	}, nil
}

var _ llm.LLM = (*LLM)(nil)

func New(dailTarget, model string, stream bool, c container.Container, callback openaiclient.Callback) llm.LLM {
	client := openaiclient.New(dailTarget, model, c, callback)
	return &LLM{
		client:   client,
		callback: callback,
		stream:   stream,
	}
}
