package models

import (
	"time"

	"gorm.io/gorm"
)

type StatusUserDB struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Status    string `json:"status"`
	UserID    int    `json:"id_user"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
