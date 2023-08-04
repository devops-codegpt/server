package models

import (
	"github.com/devops-codegpt/server/container/repository"
)

const (
	AdminRoleSort   = 0
	DefaultRoleSort = 20
	DefaultRoleId   = 2
)

type Role struct {
	Model
	Name    string  `gorm:"comment:'role name'" json:"name"`
	Keyword string  `gorm:"comment:'role keyword';unique" json:"keyword"`
	Desc    string  `gorm:"comment:'role description'" json:"desc"`
	Sort    *uint   `gorm:"default:1;comment:'Role sorting (the larger the value, the smaller the authority)'" json:"sort"`
	Status  *uint   `gorm:"comment:'role status(enable/disable)';default:1;type:tinyint(1)" json:"status"`
	Creator string  `gorm:"comment:'creator'" json:"creator"`
	Users   []*User `gorm:"foreignkey:RoleId" json:"users"`
}

func (r *Role) FindIdsBySort(rep repository.Repository, sort uint) ([]uint, error) {
	roleIds := make([]uint, 0)
	roles := make([]*Role, 0)
	err := rep.Model(&Role{}).Where("sort >= ?", sort).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}
	return roleIds, nil
}
