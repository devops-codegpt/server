package schema

import (
	"errors"
	"fmt"
	"strings"
)

// ChatMessageType is the type of chat message
type ChatMessageType string

const (
	// ChatMessageTypeAI is a message sent by an AI
	ChatMessageTypeAI ChatMessageType = "Assistant"
	// ChatMessageTypeHuman is a message sent by a user
	ChatMessageTypeHuman ChatMessageType = "Human"
	// ChatMessageTypeSystem is a message sent by the system
	ChatMessageTypeSystem ChatMessageType = "System"
)

// ChatMessage is a message sent by a user or the system.
type ChatMessage interface {
	GetText() string
	GetType() ChatMessageType
}

// Statically assert that the types implement the interface.
var (
	_ ChatMessage = AIChatMessage{}
	_ ChatMessage = HumanChatMessage{}
	_ ChatMessage = SystemChatMessage{}
)

// AIChatMessage is a message sent by an AI
type AIChatMessage struct {
	Text string
}

func (m AIChatMessage) GetType() ChatMessageType { return ChatMessageTypeAI }
func (m AIChatMessage) GetText() string          { return m.Text }

// HumanChatMessage is a message sent by a human
type HumanChatMessage struct {
	Text string
}

func (m HumanChatMessage) GetType() ChatMessageType { return ChatMessageTypeHuman }
func (m HumanChatMessage) GetText() string          { return m.Text }

// SystemChatMessage is a chat message representing information that should be instructions to the AI system
type SystemChatMessage struct {
	Text string
}

func (m SystemChatMessage) GetType() ChatMessageType { return ChatMessageTypeSystem }
func (m SystemChatMessage) GetText() string          { return m.Text }

// GetBufferString gets the buffer string of messages.
func GetBufferString(messages []ChatMessage, humanPrefix, aiPrefix string) (string, error) {
	stringMessages := make([]string, 0)
	for _, m := range messages {
		var role string
		switch m.GetType() {
		case ChatMessageTypeHuman:
			role = humanPrefix
		case ChatMessageTypeAI:
			role = aiPrefix
		case ChatMessageTypeSystem:
			role = "System"
		default:
			return "", errors.New("unexpected chat message type")
		}
		stringMessages = append(stringMessages, fmt.Sprintf("%s: %s", role, m.GetText()))
	}
	return strings.Join(stringMessages, "\n"), nil
}
