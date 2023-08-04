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
	"strconv"
	"testing"
)

var (
	roleId = uint(1)
)

func Test_roleHandler_GetRoles(t *testing.T) {
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

	mockRoleService := mock_service.NewMockRoleService(ctrl)
	mockRoleService.EXPECT().FindAllRoles(gomock.Any()).Return([]*models.Role{
		{
			Name:    "管理员",
			Keyword: "admin",
			Desc:    "",
			Sort:    &roleSort,
		},
	}, nil)

	e, cntr := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodGet, config.APIRoleList, http.NoBody)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)

	r := NewRoleHandler(cntr, mockRoleService)

	if assert.NoError(t, r.GetRoles(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_roleHandler_CreateRole(t *testing.T) {
	params := &request.RoleCreate{
		Name:    "admin",
		Keyword: "admin",
		Sort:    &roleSort,
		Desc:    "",
		Status:  nil,
		Creator: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleService := mock_service.NewMockRoleService(ctrl)
	mockRoleService.EXPECT().CreateRole(params).Return(nil)

	e, cntr := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodPost, config.APIRoleCreate, params)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)

	r := NewRoleHandler(cntr, mockRoleService)

	if assert.NoError(t, r.CreateRole(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}

func Test_roleHandler_UpdateRoleById(t *testing.T) {
	params := &request.RoleUpdate{
		Sort: &roleSort,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleService := mock_service.NewMockRoleService(ctrl)
	mockRoleService.EXPECT().UpdateById(roleId, params).Return(nil)

	e, cntr := test.PrepareForHandlerTest()
	req := test.NewJSONRequest(http.MethodPut, config.APIRoleUpdate, params)
	resp := httptest.NewRecorder()
	ctx := e.NewContext(req, resp)
	ctx.SetParamNames("roleId")
	ctx.SetParamValues(strconv.Itoa(int(roleId)))

	r := NewRoleHandler(cntr, mockRoleService)

	if assert.NoError(t, r.UpdateRoleById(ctx)) {
		assert.Equal(t, http.StatusOK, resp.Code)
	}
}
