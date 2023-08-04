package handler

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/models"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/response"
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
	roleSort = uint(1)
)

func Test_userHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := &request.UserLogin{
		Username: "admin",
		Password: "123456",
	}
	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().AuthByUserAndPassword(params.Username, params.Password).Return(true)
	mockUserService.EXPECT().FindUserByName(params.Username).Return(&models.User{
		Username: params.Username,
		ZhName:   "jack",
		Email:    "test@com",
		Role: models.Role{
			Keyword: "admin",
			Sort:    &roleSort,
		},
	}, nil).AnyTimes()

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodPost, config.APIUserLogin, params)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)
	u := NewUserHandler(container, mockUserService)

	if assert.NoError(t, u.Login(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_userHandler_GetCurrentUser(t *testing.T) {
	// Instead of the actual jwt parsing process
	saved := getJwtLoginClaims
	defer func() { getJwtLoginClaims = saved }()
	getJwtLoginClaims = func(c echo.Context) *request.JwtCustomClaims {
		return &request.JwtCustomClaims{
			UserName: "admin",
		}
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_service.NewMockUserService(ctrl)
	userInfo := &models.User{
		Username: "admin",
		ZhName:   "jack",
		Email:    "test@com",
		Role: models.Role{
			Keyword: "admin",
			Sort:    &roleSort,
		},
	}
	mockUserService.EXPECT().FindUserByName(gomock.Any()).Return(userInfo, nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodGet, config.APIUserInfo, http.NoBody)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)
	u := NewUserHandler(container, mockUserService)

	entity := newRet(Ok, OkMsg, &response.UserInfo{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		ZhName:   userInfo.ZhName,
		NickName: userInfo.Role.Name,
		RoleSort: *userInfo.Role.Sort,
	})
	if assert.NoError(t, u.GetCurrentUser(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, test.ConvertToString(entity), resp.Body.String())
	}
}

func Test_userHandler_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock_service.NewMockUserService(ctrl)
	mockUserService.EXPECT().FindAllUsers(gomock.Any()).Return([]*models.User{
		{
			Username: "admin",
			ZhName:   "jack",
			Email:    "test@com",
		},
	}, nil)

	e, container := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodGet, config.APIUserList, http.NoBody)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)
	u := NewUserHandler(container, mockUserService)

	if assert.NoError(t, u.GetUsers(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}
