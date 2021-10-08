package databases

// import (
// 	"social_media/business/users"
// 	"time"

// 	"gorm.io/gorm"
// )

// type DBRegisters struct {
// 	ID       int `gorm:"primarykey"`
// 	Nama     string
// 	Umur     int
// 	Address  string
// 	Email    string `gorm:"unique:not null"`
// 	Password string
// 	// Status    []DBStatus `gorm:"foreigenkey:IdUser"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }
// type DBStatus struct {
// 	ID        int    `gorm:"primarykey"`
// 	Status    string `gorm:"not null"`
// 	IdUser    int
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// func (user *DBRegisters) ToDomain() users.DomainRegister {
// 	return users.DomainRegister{
// 		ID:        user.ID,
// 		Name:      user.Nama,
// 		Email:     user.Email,
// 		Address:   user.Address,
// 		Password:  user.Password,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 	}
// }

// func FromDomain(domain users.DomainRegister) DBRegisters {
// 	return DBRegisters{
// 		ID:        domain.ID,
// 		Nama:      domain.Name,
// 		Email:     domain.Email,
// 		Address:   domain.Address,
// 		Password:  domain.Password,
// 		CreatedAt: domain.CreatedAt,
// 		UpdatedAt: domain.UpdatedAt,
// 	}
// }

// func (user *DBStatus) ToDomain() users.UserStatus {
// 	return users.UserStatus{
// 		ID:        user.ID,
// 		Status:    user.Status,
// 		IdUser:    user.IdUser,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 	}
// }

// func FromDomainStatus(domain *users.UserStatus) DBStatus {
// 	return DBStatus{
// 		ID:        domain.ID,
// 		Status:    domain.Status,
// 		IdUser:    domain.IdUser,
// 		CreatedAt: domain.CreatedAt,
// 		UpdatedAt: domain.UpdatedAt,
// 	}
// }

// // func (user *DBRegisters) ToDomainDetil() users.UserDeteil {
// // 	return users.UserDeteil{
// // 		DomainRegister: user.ToDomain(),
// // 		Status:  user.
// // 	}
// // }
// // func FromDomainDetiel(domain users.UserDeteil) UserDeteil {
// // 	return UserDeteil{
// // 		ID:        domain.DomainRegister.ID,
// // 		Nama:      domain.DomainRegister.Name,
// // 		Email:     domain.DomainRegister.Email,
// // 		Address:   domain.DomainRegister.Address,
// // 		Status:    domain.Status,
// // 		Password:  domain.DomainRegister.Password,
// // 		CreatedAt: domain.DomainRegister.CreatedAt,
// // 		UpdatedAt: domain.DomainRegister.UpdatedAt,
// // 	}
// // }
