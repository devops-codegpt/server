package handler

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/models"
	"github.com/devops-codegpt/server/test"
	"github.com/devops-codegpt/server/test/mocks/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_llmHandler_GetLLMs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleService := mock_service.NewMockLLMService(ctrl)
	mockRoleService.EXPECT().FindAllLLMs(gomock.Any()).Return([]*models.LLM{
		{
			Name: "chatgpt",
		},
	}, nil)

	e, cntr := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodGet, config.APILLMList, http.NoBody)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)

	l := NewLLMHandler(cntr, mockRoleService)

	if assert.NoError(t, l.GetLLMs(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}
