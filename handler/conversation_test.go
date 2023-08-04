package handler

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/models"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/test"
	"github.com/devops-codegpt/server/test/mocks/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	conversationId  = "faf5dca8-0551-4f85-9d08-289e2d424d75"
	messageId       = "faf5dca8-0551-4f85-9d08-289e2d424d79"
	parentMessageId = "faf5dca8-0551-4f85-9d08-289e2d424d76"
	username        = "admin"
)

func Test_conversationHandler_GetConversations(t *testing.T) {
	// Instead of the actual jwt parsing process
	saved := getJwtLoginClaims
	defer func() { getJwtLoginClaims = saved }()
	getJwtLoginClaims = func(c echo.Context) *request.JwtCustomClaims {
		return &request.JwtCustomClaims{
			UserName: username,
		}
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConversationService := mock_service.NewMockConversationService(ctrl)
	mockConversationService.EXPECT().FindConversations(username, gomock.Any()).Return([]*models.Conversation{
		{
			UUIDModel: models.UUIDModel{
				Id: conversationId,
			},
			Username: username,
			Title:    "",
			State:    false,
			Messages: nil,
		},
	}, nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodGet, config.APIConversationList, http.NoBody)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)
	u := NewConversationHandler(container, mockConversationService)

	if assert.NoError(t, u.GetConversations(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_conversationHandler_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConversationService := mock_service.NewMockConversationService(ctrl)
	mockConversationService.EXPECT().FindById(conversationId).Return(&models.Conversation{
		UUIDModel: models.UUIDModel{Id: conversationId},
		Username:  username,
		Title:     "",
		State:     false,
		Messages:  nil,
	}, nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodGet, config.APIConversationInfo, http.NoBody)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)
	ctx.SetParamNames("conversationId")
	ctx.SetParamValues(conversationId)
	u := NewConversationHandler(container, mockConversationService)

	if assert.NoError(t, u.GetById(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_conversationHandler_RunConversation(t *testing.T) {
	// Instead of the actual jwt parsing process
	params := &request.ConversationRun{
		ConversationId: conversationId,
		ParentMsgId:    parentMessageId,
		Role:           "Human",
		Content:        "hello",
		ContentType:    "text",
		LLM:            "chatgpt",
		Model:          "gpt-3.5-turbo",
		Stream:         true,
	}
	saved := getJwtLoginClaims
	defer func() { getJwtLoginClaims = saved }()
	getJwtLoginClaims = func(c echo.Context) *request.JwtCustomClaims {
		return &request.JwtCustomClaims{
			UserName: username,
		}
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConversationService := mock_service.NewMockConversationService(ctrl)
	mockConversationService.EXPECT().RunConversation(username, params, gomock.Any()).Return(&models.Message{
		UUIDModel:      models.UUIDModel{Id: messageId},
		Role:           "Assistant",
		Content:        "hello",
		ContentType:    "text",
		Model:          "gpt-3.5-turbo",
		Feedback:       0,
		ParentMsgId:    parentMessageId,
		ConversationId: conversationId,
	}, nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodPost, config.APIConversation, params)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)

	u := NewConversationHandler(container, mockConversationService)

	if assert.NoError(t, u.RunConversation(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_conversationHandler_RunBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := &request.BaseRun{Prompt: "hello"}
	mockConversationService := mock_service.NewMockConversationService(ctrl)
	mockConversationService.EXPECT().RunBase(params.Prompt).Return("hello", nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodPost, config.APIConversationBase, params)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)

	u := NewConversationHandler(container, mockConversationService)

	if assert.NoError(t, u.RunBase(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_conversationHandler_BatchDeleteByIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := &request.ReqDelete{Ids: "1,2,3"}
	mockConversationService := mock_service.NewMockConversationService(ctrl)
	mockConversationService.EXPECT().DeleteByIds(params.Ids).Return(nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodDelete, config.APIConversationBatchDelete, params)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)

	u := NewConversationHandler(container, mockConversationService)

	if assert.NoError(t, u.BatchDeleteByIds(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_conversationHandler_FeedbackMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := &request.MessageFeedback{
		ConversationId: conversationId,
		MessageId:      messageId,
		Feedback:       1,
	}
	mockConversationService := mock_service.NewMockConversationService(ctrl)
	mockConversationService.EXPECT().FeedbackMessage(params).Return(nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodPost, config.APIMessageFeedback, params)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)

	u := NewConversationHandler(container, mockConversationService)

	if assert.NoError(t, u.FeedbackMessage(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}
