package userDB

import (
	// "social_media/business/users"
	"social_media/business/users/userRegister"
	"social_media/drivers/databases/userStatus"
	"time"

	"gorm.io/gorm"
)

type DBRegisters struct {
	ID        int `gorm:"primarykey"`
	Nama      string
	Umur      int
	Address   string
	Email     string `gorm:"unique:not null"`
	Password  string
	DBStatus  []userStatus.DBStatus `gorm:"foreigenKey:DBRegistersID ;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *DBRegisters) ToDomain() userRegister.DomainRegister {
	return userRegister.DomainRegister{
		ID:        user.ID,
		Name:      user.Nama,
		Email:     user.Email,
		Address:   user.Address,
		Password:  user.Password,
		Status:    user.DBStatus,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain userRegister.DomainRegister) DBRegisters {
	return DBRegisters{
		ID:        domain.ID,
		Nama:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Password:  domain.Password,
		DBStatus:  domain.Status,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func ToListDeteil(data []DBRegisters) (result []userRegister.DomainRegister) {
	result = []userRegister.DomainRegister{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
