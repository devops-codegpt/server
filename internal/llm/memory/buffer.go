package memory

import (
	"errors"
	"fmt"
	"github.com/devops-codegpt/server/internal/llm/schema"
)

// Buffer is a simple form of memory that remembers previous conversational back and forth directly.
type Buffer struct {
	ChatHistory *ChatMessageHistory

	ReturnMessages bool
	InputKey       string
	OutputKey      string
	HumanPrefix    string
	AIPrefix       string
	MemoryKey      string
	Depth          int
}

// Statically assert that Buffer implement the memory interface.
var _ schema.Memory = &Buffer{}

// NewBuffer is a function for crating a new buffer memory.
func NewBuffer() *Buffer {
	m := Buffer{
		ChatHistory:    NewChatMessageHistory(),
		ReturnMessages: false,
		InputKey:       "",
		OutputKey:      "",
		HumanPrefix:    "Human",
		AIPrefix:       "Assistant",
		MemoryKey:      "history",
		Depth:          2, // set a depth=2, to only keep the last 2 interactions in memory
	}

	return &m
}

// MemoryVariables gets the input key the buffer memory class will load dynamically.
func (m *Buffer) MemoryVariables() []string {
	return []string{m.MemoryKey}
}

// LoadMemoryVariables returns the previous chat messages stored in memory. Previous chat messages
// are returned in a map with the key specified in the MemoryKey field. This key defaults to
// "history". If ReturnMessages is set to true the output is a slice of schema.ChatMessage. Otherwise,
// the output is a buffer string of the chat messages.
func (m *Buffer) LoadMemoryVariables(map[string]any) (map[string]any, error) {
	m.ChatHistory.Keep(m.Depth)
	if m.ReturnMessages {
		return map[string]any{
			m.MemoryKey: m.ChatHistory.messages,
		}, nil
	}

	bufferString, err := schema.GetBufferString(m.ChatHistory.messages, m.HumanPrefix, m.AIPrefix)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		m.MemoryKey: bufferString,
	}, nil
}

// SaveContext uses the input values to the llm to save a user message, and the output values
// of the llm to save an AI message. If the input or output key is not set, the input values or
// output values must contain only one key such that the function can know what string to
// add as a user and AI message. On the other hand, if the output key or input key is set, the
// input key must be a key in the input values and the output key must be a key in the output
// values. The values in the input and output values used to save a user and AI message must
// be strings.
func (m *Buffer) SaveContext(inputValues, outputValues map[string]any) error {
	userInputValue, err := getInputValue(inputValues, m.InputKey)
	if err != nil {
		return err
	}

	m.ChatHistory.AddUserMessage(userInputValue)

	aiOutputValue, err := getInputValue(outputValues, m.OutputKey)
	if err != nil {
		return err
	}

	m.ChatHistory.AddAIMessage(aiOutputValue)

	return nil
}

// Clear sets the chat messages to a new and empty chat message history.
func (m *Buffer) Clear() error {
	m.ChatHistory.Clear()
	return nil
}

// ErrInvalidInputValues is returned when input values given to a memory in save context are invalid.
var ErrInvalidInputValues = errors.New("invalid input values")

func getInputValue(inputValues map[string]any, inputKey string) (string, error) {
	// If the input key is set, return the value in the inputValues with the input key.
	if inputKey != "" {
		inputValue, ok := inputValues[inputKey]
		if !ok {
			return "", fmt.Errorf(
				"%w: %v do not contain inputKey %s",
				ErrInvalidInputValues,
				inputValues,
				inputKey,
			)
		}

		return getInputValueReturnToString(inputValue)
	}

	// Otherwise error if length of map isn't one, or return the only entry in the map.
	if len(inputValues) > 1 {
		return "", fmt.Errorf(
			"%w: multiple keys and no input key set",
			ErrInvalidInputValues,
		)
	}

	for _, inputValue := range inputValues {
		return getInputValueReturnToString(inputValue)
	}

	return "", fmt.Errorf("%w: 0 keys", ErrInvalidInputValues)
}

func getInputValueReturnToString(
	inputValue any,
) (string, error) {
	switch value := inputValue.(type) {
	case string:
		return value, nil
	default:
		return "", fmt.Errorf(
			"%w: input value %v not string",
			ErrInvalidInputValues,
			inputValue,
		)
	}
}
