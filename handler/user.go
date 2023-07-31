package handler

import (
	"errors"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/response"
	"github.com/devops-codegpt/server/service"
	"github.com/devops-codegpt/server/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserHandler interface {
	Login(c echo.Context) error
	GetCurrentUser(c echo.Context) error
	GetUsers(c echo.Context) error
}

type userHandler struct {
	context container.Container
	service service.UserService
}

// Login checks username and password, return jwt
func (u *userHandler) Login(c echo.Context) error {
	var req request.UserLogin
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	authenticate := u.service.AuthByUserAndPassword(req.Username, req.Password)
	if !authenticate {
		return c.NoContent(http.StatusUnauthorized)
	}

	// Add user if user not in repo
	_, err := u.service.FindUserByName(req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		rc := request.UserCreate{
			Username: req.Username,
		}
		err = u.service.CreateUser(&rc)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}
	}

	// Get user
	user, err := u.service.FindUserByName(req.Username)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	config := u.context.GetConfig()
	// Construct JWT and return it
	claims := &request.JwtCustomClaims{
		UserName:    user.Username,
		RoleKeyword: user.Role.Keyword,
		RoleSort:    *user.Role.Sort,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.Auth.Expires))),
		},
	}
	// Create claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as ret
	t, err := token.SignedString([]byte(config.Auth.JwtKey))
	if err != nil {
		return failWithCode(c, http.StatusBadRequest)
	}
	return successWithData(c, echo.Map{
		"token": t,
	})
}

func (u *userHandler) GetCurrentUser(c echo.Context) error {
	claims := getJwtLoginClaims(c)
	username := claims.UserName
	user, err := u.service.FindUserByName(username)
	if err != nil {
		return failWithMsg(c, err.Error())
	}

	resp := response.UserInfo{}
	// Hide some response fields
	utils.Struct2Struct(user, &resp)
	resp.NickName = user.Role.Name
	resp.RoleSort = *user.Role.Sort
	return successWithData(c, resp)
}

func (u *userHandler) GetUsers(c echo.Context) error {
	var req request.UserList
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, req)
	}
	users, err := u.service.FindAllUsers(&req)
	if err != nil {
		return failWithMsg(c, err.Error())
	}

	// Convert to ResponseStruct, Hide some no used fields.
	userList := make([]*response.UserList, 0)
	utils.Struct2Struct(users, &userList)
	var resp response.PageData
	resp.PageInfo = req.PageInfo
	resp.List = userList
	return successWithData(c, resp)
}

// getJwtLoginClaims gets the currently logged-in user's jwt
var getJwtLoginClaims = func(c echo.Context) *request.JwtCustomClaims {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*request.JwtCustomClaims)
	return claims
}

func NewUserHandler(c container.Container, s service.UserService) UserHandler {
	if s == nil {
		s = service.NewUserService(c)
	}
	return &userHandler{
		context: c,
		service: s,
	}
}
