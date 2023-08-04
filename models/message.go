package models

import (
	"github.com/devops-codegpt/server/container/repository"
)

const (
	DefaultMessageUserRole = "Human"
	DefaultMessageBotRole  = "Assistant"
	DefaultBotContentType  = "text"
	// DefaultHistoryDepth is a complete user-bot log
	DefaultHistoryDepth = 2
)

// Message Store conversation logs, Including user questions and chat bot responses
type Message struct {
	UUIDModel
	Role           string       `gorm:"comment:'informative role'" json:"role"`
	Content        string       `gorm:"type:text;" json:"content"`
	ContentType    string       `gorm:"comment:'content type'" json:"contentType"`
	Model          string       `gorm:"comment:'ai models name'" json:"models"`
	Feedback       uint         `gorm:"type:tinyint(1)" json:"feedback"`
	ParentMsgId    string       `gorm:"type:varchar(36)" json:"parentMsgId"`
	ConversationId string       `gorm:"type:varchar(36)" json:"conversationId"`
	Conversation   Conversation `gorm:"foreignkey:ConversationId" json:"conversation"`
}

type History struct {
	Role    string
	Content string
}

// FindPreMessages Find previous message records by depth. The depth is an even number, which means a completed user-bot log
func (m *Message) FindPreMessages(rep repository.Repository, depth uint) ([]*History, error) {
	histories := make([]*History, 0)
	if depth == 0 || m.ParentMsgId == "" {
		return histories, nil
	}
	preMessage := &Message{}
	err := rep.Where("id = ?", m.ParentMsgId).First(preMessage).Error
	if err != nil {
		return nil, err
	}
	histories = append(histories, &History{Role: preMessage.Role, Content: preMessage.Content})

	// Get previous histories
	preHistories, err := preMessage.FindPreMessages(rep, depth-1)
	if err != nil {
		return nil, err
	}
	histories = append(histories, preHistories...)
	return histories, nil
}
