package models

import (
	"github.com/devops-codegpt/server/container/repository"
)

const (
	UserStatusNormal   uint = 1
	UserStatusDisabled uint = 0
)

type User struct {
	Model
	Username string `gorm:"comment:'user name';unique" json:"username"`
	ZhName   string `gorm:"comment:'user zh name'" json:"zhName"`
	Email    string `gorm:"comment:'user email'" json:"email"`
	Status   *uint  `gorm:"comment:'user status';default:1;type:tinyint(1)" json:"status"`
	Creator  string `gorm:"comment:'creator'" json:"creator"`
	RoleId   uint   `gorm:"comment:'role id'" json:"roleId"`
	Role     Role   `gorm:"foreignkey:RoleId" json:"role"`
}

func (u *User) FindByName(rep repository.Repository, name string) (*User, error) {
	var user *User

	err := rep.Preload("role").Where("username = ?", name).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
