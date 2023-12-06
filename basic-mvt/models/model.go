package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `gorm:"primarykey" json:"id" query:"id" form:"id"`
	CreatedAt time.Time      `json:"created_at" query:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" query:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" query:"deleted_at" form:"deleted_at"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

type Error struct {
	Message string `json:"message"`
}
