package request

import (
	"github.com/devops-codegpt/server/models"
)

type ConversationList struct {
	models.PageInfo
}

type ConversationRun struct {
	ConversationId string `json:"conversationId" form:"conversationId" validate:"required,uuid"`
	ParentMsgId    string `json:"parentMsgId,omitempty" form:"parentMsgId"`
	Role           string `json:"role" form:"role" validate:"required"`
	Content        string `json:"content" form:"content" validate:"required"`
	ContentType    string `json:"contentType" form:"contentType" validate:"required"`
	LLM            string `json:"llm" form:"llm" validate:"required"`
	Model          string `json:"models" form:"models"`
	Stream         bool   `json:"stream" form:"stream" validate:"required"`
}

type BaseRun struct {
	Prompt string `json:"prompt" form:"prompt" validate:"required"`
}

type MessageFeedback struct {
	ConversationId string `json:"conversationId" form:"conversationId" validate:"required,uuid"`
	MessageId      string `json:"messageId" form:"conversationId" validate:"required,uuid"`
	Feedback       uint   `json:"feedback" form:"feedback" validate:"required"`
}

type ConversationWS struct {
	Command string `json:"command"`
	Prompt  string `json:"prompt"`
}
