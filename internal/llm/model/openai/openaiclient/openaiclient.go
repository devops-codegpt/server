package openaiclient

import (
	"context"
	"github.com/devops-codegpt/server/container"
)

type Client struct {
	dailTarget string
	model      string
	container  container.Container
	callback   Callback
}

// Callback handlers for streaming. Only works with LLMs that support streaming.
type Callback interface {
	// OnLLMNewToken Runs on new LLM token. Only available when streaming is enabled
	OnLLMNewToken(str string)
}

// New returns a new OpenAI client.
func New(dailTarget, model string, c container.Container, cb Callback) *Client {
	return &Client{
		dailTarget: dailTarget,
		model:      model,
		container:  c,
		callback:   cb,
	}
}

// CreateChatWithStream creates chat request
func (c *Client) CreateChatWithStream(ctx context.Context, r *ChatRequest) (string, error) {
	r.Model = c.model
	if r.Model == "" {
		r.Model = defaultChatModel
	}
	resp, err := c.createChatWithStream(ctx, r)
	if err != nil {
		return "", err
	}
	return resp, nil
}
