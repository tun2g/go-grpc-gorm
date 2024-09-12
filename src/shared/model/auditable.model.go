package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditableModel struct {
	Id        string         `gorm:"type:varchar(36);primary_key;" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deletedAt"`
}

func (base *AuditableModel) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	base.Id = uuid
	return nil
}
