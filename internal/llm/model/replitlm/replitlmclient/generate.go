package replitlmclient

import (
	"context"
	"errors"
	"github.com/devops-codegpt/server/internal/llm/proto/replitlmpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GenerateRequest struct {
	Prompt string `json:"prompt"`
	Lang   string `json:"lang"`
}

type GenerateResponse struct {
	CodeList         []string `json:"codeList"`
	CompletionTokens int      `json:"completionTokens"`
	PromptTokens     int      `json:"promptTokens"`
}

func (c *Client) runGenerate(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error) {
	conn, err := grpc.Dial(c.dailTarget,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		e := conn.Close()
		if e != nil {
			c.container.GetLogger().GetZapLogger().Errorf("close grpc service failed")
		}
	}(conn)

	client := replitlmpb.NewCodeGeneratorClient(conn)
	request := &replitlmpb.CodeRequest{
		Prompt: req.Prompt,
		Lang:   req.Lang,
	}
	response, err := client.SendPrompt(ctx, request)
	if err != nil {
		c.container.GetLogger().GetZapLogger().Errorf("failed to call code-generate service failed: %v", err)
		return nil, err
	}
	if response.Code != 0 {
		c.container.GetLogger().GetZapLogger().Errorf("code generate failed: %s", response.Msg)
		return nil, errors.New(response.Msg)
	}
	return &GenerateResponse{
		CodeList:         response.Ret.CodeList,
		CompletionTokens: int(response.Ret.CompletionTokenNum),
		PromptTokens:     int(response.Ret.PromptTokenNum),
	}, nil
}
