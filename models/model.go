package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at" json:"deletedAt"`
}

type UUIDModel struct {
	Id        string         `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at" json:"deletedAt"`
}

func (u *UUIDModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}

type DomainObject interface {
	User | Role | Conversation | LLM
}

func FindList[T DomainObject](db *gorm.DB, pageInfo *PageInfo, list []*T) ([]*T, error) {
	err := db.Count(&pageInfo.Total).Error
	if err != nil {
		return nil, err
	}
	if pageInfo.NoPagination {
		err = db.Find(&list).Error
	} else {
		limit, offset := pageInfo.GetLimit()
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}
	return list, err
}
