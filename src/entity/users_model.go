package entity

import (
	"time"

	"gorm.io/gorm"
)

type UsersModel struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     *string        `json:"email"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime" `
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
