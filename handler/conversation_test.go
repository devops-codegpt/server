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

var (
	conversationId = "faf5dca8-0551-4f85-9d08-289e2d424d75"
)

func Test_conversationHandler_GetConversations(t *testing.T) {
	// Instead of the actual jwt parsing process
	username := "admin"
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
		Username:  "admin",
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
