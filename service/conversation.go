package service

import (
	"context"
	"fmt"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/internal/llm/memory"
	"github.com/devops-codegpt/server/internal/llm/model/openai"
	oc "github.com/devops-codegpt/server/internal/llm/model/openai/openaiclient"
	"github.com/devops-codegpt/server/models"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/service/discovery"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

const (
	chatgpt   = "chatgpt"
	chatModel = "gpt-3.5-turbo"
)

type ConversationService interface {
	FindConversations(username string, req *request.ConversationList) ([]*models.Conversation, error)
	FindById(id string) (*models.Conversation, error)
	RunConversation(username string, req *request.ConversationRun, callback oc.Callback) (*models.Message, error)
	RunConversationWS(command, prompt string, buffer *memory.Buffer, callback oc.Callback) error
	RunBase(prompt string) (string, error)
	DeleteByIds(ids []uuid.UUID) error
	FeedbackMessage(req *request.MessageFeedback) error
}

type conversationService struct {
	container container.Container
}

// FindConversations gets all conversations for a user
func (cs *conversationService) FindConversations(username string, req *request.ConversationList) ([]*models.Conversation, error) {
	rep := cs.container.GetRepository()
	conversations := make([]*models.Conversation, 0)
	db := rep.Model(&models.Conversation{}).Order("created_at DESC")
	db = db.Where("username = ?", username)
	return models.FindList(db, &req.PageInfo, conversations)
}

func (cs *conversationService) FindById(id string) (*models.Conversation, error) {
	conversationId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	conversation := &models.Conversation{}
	rep := cs.container.GetRepository()
	err = rep.Preload("Messages").Where("id = ?", conversationId).First(conversation).Error
	if err != nil {
		return nil, err
	}
	return conversation, nil
}

func (cs *conversationService) RunConversation(name string, req *request.ConversationRun, cb oc.Callback) (*models.Message, error) {
	rep := cs.container.GetRepository()
	conversation := &models.Conversation{}

	err := rep.Model(&models.Conversation{}).Where("id = ?", req.ConversationId).First(conversation).Error
	if err == gorm.ErrRecordNotFound {
		// Create a new conversation
		conversation.Id = req.ConversationId
		conversation.Username = name
		if e := rep.Create(conversation).Error; e != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	// Add user message to database
	userMessage := &models.Message{
		UUIDModel:      models.UUIDModel{Id: uuid.New().String()},
		Role:           models.DefaultMessageUserRole,
		Content:        req.Content,
		ContentType:    req.ContentType,
		Model:          req.Model,
		ParentMsgId:    req.ParentMsgId,
		ConversationId: req.ConversationId,
	}
	err = rep.Model(conversation).Association("Messages").Append(userMessage)
	if err != nil {
		return nil, err
	}
	// Get llm service from service discover center
	discover := discovery.New(cs.container)
	openaiGrpc, err := discover.GetServiceByName(req.LLM)
	if err != nil {
		return nil, err
	}
	llmClient := openai.New(
		fmt.Sprintf("%s:%d", openaiGrpc.Address, openaiGrpc.Port),
		req.Model,
		req.Stream,
		cs.container,
		cb,
	)
	// Construct prompt
	histories, err := userMessage.FindPreMessages(rep, models.DefaultHistoryDepth)
	if err != nil {
		return nil, err
	}
	prompt := constructHistoryStr(histories) + fmt.Sprintf(
		"\n%s: %s\n%s:",
		userMessage.Role,
		userMessage.Content,
		models.DefaultMessageBotRole,
	)
	answer, err := llmClient.Call(context.Background(), prompt)
	if err != nil {
		return nil, err
	}
	botMessage := &models.Message{
		UUIDModel:      models.UUIDModel{Id: uuid.New().String()},
		Role:           models.DefaultMessageBotRole,
		Content:        answer,
		ContentType:    models.DefaultBotContentType,
		Model:          req.Model,
		ParentMsgId:    userMessage.Id,
		ConversationId: req.ConversationId,
	}
	// Save bot message
	err = rep.Model(conversation).Association("Messages").Append(botMessage)
	if err != nil {
		return nil, err
	}
	return botMessage, nil
}

func (cs *conversationService) FeedbackMessage(req *request.MessageFeedback) error {
	rep := cs.container.GetRepository()
	err := rep.Model(&models.Message{}).
		Where("id = ? AND conversation_id = ?", req.MessageId, req.ConversationId).
		Update("feedback", req.Feedback).Error
	return err
}

func (cs *conversationService) DeleteByIds(ids []uuid.UUID) error {
	rep := cs.container.GetRepository()

	// Delete conversation-related messages
	err := rep.Where("conversation_id IN (?)", ids).Delete(&models.Message{}).Error
	if err != nil {
		return err
	}
	return rep.Where("id IN (?)", ids).Delete(&models.Conversation{}).Error
}

func (cs *conversationService) RunConversationWS(command, prompt string, buffer *memory.Buffer, callback oc.Callback) error {
	discover := discovery.New(cs.container)
	openaiGrpc, err := discover.GetServiceByName(command)
	if err != nil {
		return err
	}
	llmClient := openai.New(
		fmt.Sprintf("%s:%d", openaiGrpc.Address, openaiGrpc.Port),
		chatModel,
		true,
		cs.container,
		callback,
	)
	history, err := buffer.LoadMemoryVariables(map[string]any{})
	if err != nil {
		return err
	}
	totalPrompt := prompt
	if command == chatgpt {
		totalPrompt = history["history"].(string) + fmt.Sprintf("\nHuman: %s\n", prompt)
	}
	answer, err := llmClient.Call(context.Background(), totalPrompt)
	if err != nil {
		return err
	}
	// Add to buffer.
	err = buffer.SaveContext(map[string]any{"Human": prompt}, map[string]any{"Assistant": answer})
	if err != nil {
		return err
	}
	return nil
}

func (cs *conversationService) RunBase(prompt string) (string, error) {
	command := chatgpt
	discover := discovery.New(cs.container)
	openaiGrpc, err := discover.GetServiceByName(command)
	if err != nil {
		return "", err
	}
	llmClient := openai.New(
		fmt.Sprintf("%s:%d", openaiGrpc.Address, openaiGrpc.Port),
		chatModel,
		false,
		cs.container,
		nil,
	)
	answer, err := llmClient.Call(context.Background(), prompt)
	if err != nil {
		return "", err
	}
	return answer, nil
}

func constructHistoryStr(histories []*models.History) string {
	historyStrs := make([]string, 0)
	// The head history is the most recent record. Histories are like stack
	if len(histories) == 0 {
		return ""
	}
	for i := len(histories) - 1; i > 0; i++ {
		historyStrs = append(historyStrs, fmt.Sprintf("%s: %s", histories[i].Role, histories[i].Content))
	}
	return strings.Join(historyStrs, "\n")
}

func NewConversationService(c container.Container) ConversationService {
	return &conversationService{container: c}
}
