package request

import "github.com/devops-codegpt/server/models"

type RoleCreate struct {
	Name    string `json:"name" form:"name" validate:"required"`
	Keyword string `json:"keyword" form:"keyword" validate:"required"`
	Sort    *uint  `json:"sort" form:"sort" validate:"required"`
	Desc    string `json:"desc" form:"desc"`
	Status  *uint  `json:"status" form:"status"`
	Creator string `json:"creator" form:"creator"`
}

type RoleUpdate struct {
	Name    string `json:"name"`
	Keyword string `json:"keyword"`
	Sort    *uint  `json:"sort"`
	Desc    string `json:"desc"`
	Status  uint   `json:"status"`
	Creator string `json:"creator"`
}

type RoleList struct {
	Name            string `json:"name" form:"name"`
	Keyword         string `json:"keyword" form:"keyword"`
	Status          *uint  `json:"status" form:"status"`
	CurrentRoleSort uint   `json:"currentRoleSort"`
	models.PageInfo
}
