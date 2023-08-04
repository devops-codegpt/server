package service

import (
	"fmt"
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/models"
	"github.com/devops-codegpt/server/request"
	"github.com/devops-codegpt/server/utils"
	"strings"
)

type RoleService interface {
	FindAllRoles(req *request.RoleList) ([]*models.Role, error)
	CreateRole(req *request.RoleCreate) error
	UpdateById(id uint, req *request.RoleUpdate) error
}

type roleService struct {
	container container.Container
}

func (r *roleService) FindAllRoles(req *request.RoleList) ([]*models.Role, error) {
	rep := r.container.GetRepository()

	roles := make([]*models.Role, 0)
	db := rep.Model(&models.Role{}).Order("created_at DESC")
	// Only get characters that are greater than the current character sort
	db = db.Where("sort >= ?", req.CurrentRoleSort)

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	keyword := strings.TrimSpace(req.Keyword)
	if keyword != "" {
		db = db.Where("keyword LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword))
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
	return models.FindList(db, &req.PageInfo, roles)
}

func (r *roleService) CreateRole(req *request.RoleCreate) error {
	rep := r.container.GetRepository()
	role := &models.Role{}
	utils.Struct2Struct(req, role)
	return rep.Create(role).Error
}

func (r *roleService) UpdateById(id uint, req *request.RoleUpdate) error {
	rep := r.container.GetRepository()
	return rep.Model(&models.Role{}).Where("id = ?", id).Updates(req).Error
}

func NewRoleService(c container.Container) RoleService {
	return &roleService{container: c}
}
