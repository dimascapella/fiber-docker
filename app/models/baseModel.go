package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int            `gorm:"primary_key;not null;autoIncrement" json:"id"  validate:"uuid,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
