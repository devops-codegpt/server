package llm

import "context"

// LLM is a Large Language Model.
type LLM interface {
	Call(ctx context.Context, prompt string, options ...CallOption) (string, error)
	Generate(ctx context.Context, prompts []string, options ...CallOption) (*AIResponse, error)
}

type AIResponse struct {
	Choices []*Choice `json:"choices"`
	Usage   AIUsage   `json:"usage"`
}

type Choice struct {
	// Text is the generated text.
	Text string `json:"text"`
	// Info is the generation info, contains specific information.
	Info map[string]any `json:"info"`
}

type AIUsage struct {
	PromptTokens     int `json:"promptTokens"`
	CompletionTokens int `json:"completionTokens"`
	TotalTokens      int `json:"totalTokens"`
}
