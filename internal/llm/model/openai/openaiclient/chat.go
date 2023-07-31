package openaiclient

import (
	"context"
	"errors"
	"github.com/devops-codegpt/server/internal/llm/proto/openaipb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

const (
	defaultChatModel = "gpt-3.5-turbo"
)

// ChatRequest is a request to create an embedding.
type ChatRequest struct {
	Model       string  `json:"models"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature,omitempty"`
}

// ChatMessage is a message in a chat request.
type ChatMessage struct {
	// The role of the author of this message. One of system, user, or assistant.
	Role string `json:"role"`
	// The content of the message.
	Content string `json:"content"`
	// The name of the author of this message. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name string `json:"name,omitempty"`
}

func (c *Client) createChatWithStream(ctx context.Context, message *ChatRequest) (string, error) {
	// Connect openai GRPC service
	conn, err := grpc.Dial(c.dailTarget,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer func(conn *grpc.ClientConn) {
		e := conn.Close()
		if e != nil {
			c.container.GetLogger().GetZapLogger().Errorf("close grpc service failed")
		}
	}(conn)

	client := openaipb.NewChatgptClient(conn)
	// Construct grpc request
	req := &openaipb.Message{}
	req.Content = message.Prompt
	stream, err := client.Send(ctx, req)
	if err != nil {
		c.container.GetLogger().GetZapLogger().Errorf("failed to call server-side streaming method: %v", err)
		return "", err
	}
	resp := ""
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// After the server has sent all the responses, it will return EOF
			break
		}
		if err != nil {
			return "", err
		}

		// Handle response
		if outputErr := response.GetError(); outputErr != "" {
			return "", errors.New(outputErr)
		}
		if c.callback != nil {
			c.callback.OnLLMNewToken(response.GetContent())
		}
		resp += response.GetContent()
	}
	return resp, nil
}
