package models

import (
	"time"

	"gorm.io/gorm"
)

type UserDatabase struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	Nama          string         `json:"nama"`
	Umur          int            `json:"umur"`
	Alamat        string         `json:"alamat"`
	Email         string         `json:"email" gorm:"unique:not null"`
	Password      string         `json:"password"`
	StatusUserDBs []StatusUserDB `gorm:"foreigenKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
