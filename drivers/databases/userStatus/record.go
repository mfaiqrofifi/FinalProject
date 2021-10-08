package userStatus

import (
	"social_media/business/users/userStatus"
	"social_media/drivers/databases/detailStatus"
	"time"

	"gorm.io/gorm"
)

type DBStatus struct {
	ID            int    `gorm:"primarykey"`
	Status        string `gorm:"not null"`
	DBRegistersID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Coment        []detailStatus.DetailStatusDB `gorm:"foreignKey:DBStatusID;references:ID"`
	DeletedAt     gorm.DeletedAt                `gorm:"index"`
}

func (user *DBStatus) ToDomain() userStatus.UserStatus {
	return userStatus.UserStatus{
		ID:        user.ID,
		Status:    user.Status,
		IdUser:    user.DBRegistersID,
		Comment:   user.Coment,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomainStatus(domain *userStatus.UserStatus) DBStatus {
	return DBStatus{
		ID:            domain.ID,
		Status:        domain.Status,
		DBRegistersID: domain.IdUser,
		Coment:        domain.Comment,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func ToListDeteil(data []DBStatus) (result []userStatus.UserStatus) {
	result = []userStatus.UserStatus{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
