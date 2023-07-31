package handler

import (
	"fmt"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/response"
	"github.com/devops-codegpt/server/service"
	"github.com/devops-codegpt/server/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoleHandler interface {
	GetRoles(c echo.Context) error
	CreateRole(c echo.Context) error
	UpdateRoleById(c echo.Context) error
}

type roleHandler struct {
	context container.Container
	service service.RoleService
}

func (r *roleHandler) GetRoles(c echo.Context) error {
	var req request.RoleList
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	// Bind current user role sorting
	jwtClaims := getJwtLoginClaims(c)
	req.CurrentRoleSort = jwtClaims.RoleSort

	roles, err := r.service.FindAllRoles(&req)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	//  Hide some fields
	list := make([]*response.RoleList, 0)
	utils.Struct2Struct(roles, &list)
	var resp response.PageData
	resp.PageInfo = req.PageInfo
	resp.List = list
	return successWithData(c, resp)
}

func (r *roleHandler) CreateRole(c echo.Context) error {
	var req request.RoleCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	// Param validate
	if err := c.Validate(req); err != nil {
		return err
	}

	err := r.service.CreateRole(&req)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	return success(c)
}

func (r *roleHandler) UpdateRoleById(c echo.Context) error {
	var req request.RoleUpdate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	fmt.Printf("role: %+v", req)
	roleId := utils.Str2Uint(c.Param("roleId"))
	err := r.service.UpdateById(roleId, &req)
	if err != nil {
		return failWithMsg(c, err.Error())
	}
	return success(c)
}

func NewRoleHandler(c container.Container) RoleHandler {
	return &roleHandler{
		context: c,
		service: service.NewRoleService(c),
	}
}
