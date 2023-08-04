package request

import (
	"github.com/devops-codegpt/server/models"
	"github.com/golang-jwt/jwt/v4"
)

// JwtCustomClaims are custom claims extending default ones
type JwtCustomClaims struct {
	UserName    string `json:"userName"`
	RoleKeyword string `json:"roleKeyword"`
	RoleSort    uint   `json:"roleSort"`
	jwt.RegisteredClaims
}

// UserLogin represents user authentication request body
type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserList struct {
	CurrentRoleSort uint
	Username        string `json:"username" form:"username"`
	ZhName          string `json:"zhName" form:"zhName"`
	Email           string `json:"email" form:"email"`
	Status          *uint  `json:"status" form:"status"`
	models.PageInfo
}

type UserCreate struct {
	Username string `json:"username" form:"username" validate:"required"`
	ZhName   string `json:"zhName" form:"zhName" validate:"required"`
	Email    string `json:"email" form:"email"`
	Status   uint   `json:"status,omitempty" form:"status"`
	RoleId   uint   `json:"roleId" form:"roleId" validate:"required"`
}

// FieldTrans translates the name of the field that needs to be verified
func (c *UserCreate) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "user"
	m["ZhName"] = "user Chinese name"
	m["RoleId"] = "ID of the role the user belongs to"
	return m
}

type UserUpdate struct {
	Status *uint `json:"status"`
	RoleId uint  `json:"roleId" validate:"required"`
}
