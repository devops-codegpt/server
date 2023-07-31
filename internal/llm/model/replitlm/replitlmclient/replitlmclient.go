package replitlmclient

import (
	"context"
	"github.com/devops-codegpt/server/container"
)

type Client struct {
	dailTarget string
	container  container.Container
}

func (c *Client) RunGenerate(ctx context.Context, r *GenerateRequest) (*GenerateResponse, error) {
	return c.runGenerate(ctx, r)
}

// New returns a new OpenAI client.
func New(dailTarget string, ctr container.Container) *Client {
	return &Client{
		dailTarget: dailTarget,
		container:  ctr,
	}
}
