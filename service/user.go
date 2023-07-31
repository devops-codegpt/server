package service

import (
	"fmt"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/models"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/utils"
	"strings"
)

type UserService interface {
	AuthByUserAndPassword(username, password string) bool
	FindUserByName(username string) (*models.User, error)
	FindAllUsers(req *request.UserList) ([]*models.User, error)
	CreateUser(req *request.UserCreate) error
}

type userService struct {
	container container.Container
}

func (u *userService) AuthByUserAndPassword(username, password string) bool {
	logger := u.container.GetLogger()
	// Check username and password
	if username != "admin" || password != "123456" {
		logger.GetZapLogger().Errorf("wrong username： %s or password: %v", username, password)
		return false
	}
	return true
}

func (u *userService) FindUserByName(username string) (*models.User, error) {
	rep := u.container.GetRepository()

	// Get user from database
	user := &models.User{}
	err := rep.Preload("Role").Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) FindAllUsers(req *request.UserList) ([]*models.User, error) {
	rep := u.container.GetRepository()

	users := make([]*models.User, 0)
	db := rep.Model(&models.User{}).Order("created_at DESC")
	// Determine whether it is a super administrator
	if req.CurrentRoleSort != models.AdminRoleSort {
		role := &models.Role{}
		roleIds, err := role.FindIdsBySort(rep, req.CurrentRoleSort)
		if err != nil {
			return nil, err
		}
		db = db.Where("role_id in (?)", roleIds)
	}

	username := strings.TrimSpace(req.Username)
	if username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	zhName := strings.TrimSpace(req.ZhName)
	if zhName != "" {
		db = db.Where("zh_name LIKE ?", fmt.Sprintf("%%%s%%", req.ZhName))
	}
	email := strings.TrimSpace(req.Email)
	if email != "" {
		db = db.Where("email LIKE ?", fmt.Sprintf("%%%s%%", req.Email))
	}
	if req.Status != nil {
		db = db.Where("status = ?", req.Status)
	}

	// Calculate the total number of queries
	err := db.Count(&req.PageInfo.Total).Error
	if err != nil {
		return nil, err
	}

	// Calculate pagination based on page number
	return models.FindList(db, &req.PageInfo, users)
}

// CreateUser Adds user information to the database
func (u *userService) CreateUser(req *request.UserCreate) error {
	req.Email = "example@test.com"
	req.ZhName = "test"

	var user models.User
	utils.Struct2Struct(req, &user)
	user.Creator = "auto create"
	user.RoleId = models.DefaultRoleId

	rep := u.container.GetRepository()
	return rep.Model(&models.User{}).Create(&user).Error
}

func NewUserService(c container.Container) UserService {
	return &userService{container: c}
}
