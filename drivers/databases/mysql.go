package databases

// import (
// 	"context"
// 	"social_media/business/users"

// 	"gorm.io/gorm"
// )

// type MysqlUserRepository struct {
// 	Conn *gorm.DB
// }

// func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
// 	return &MysqlUserRepository{
// 		Conn: conn,
// 	}
// }

// func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.DomainRegister, error) {
// 	var user DBRegisters
// 	result := rep.Conn.First(&user, "email=?", email)
// 	if result.Error != nil {
// 		return users.DomainRegister{}, result.Error
// 	}
// 	return user.ToDomain(), nil

// }
// func (rep *MysqlUserRepository) Register(ctx context.Context, name string, email string, address string, password string) (users.DomainRegister, error) {
// 	var user DBRegisters
// 	user.Nama = name
// 	user.Address = address
// 	user.Email = email
// 	user.Password = password
// 	result := rep.Conn.Create(&user)
// 	if result.Error != nil {
// 		return users.DomainRegister{}, result.Error
// 	}
// 	return user.ToDomain(), nil

// }

// func (rep *MysqlUserRepository) CreateStatus(ctx context.Context, status string, IdUser int) (users.UserStatus, error) {
// 	var user DBStatus
// 	user.Status = status
// 	user.IdUser = IdUser
// 	result := rep.Conn.Create(&user)
// 	if result.Error != nil {
// 		return users.UserStatus{}, result.Error
// 	}
// 	return user.ToDomain(), nil

// }

// // func (rep *MysqlUserRepository) Detiled(ctx context.Context, nama string, alamat string,status DBStatus) (users.DomainRegister, error) {
// // 	var user [] DBRegisters
// // 	result := rep.Conn.Preload("status").Find(&user)
// // 	if result.Error != nil {
// // 		return users.DomainRegister{}, result.Error
// // 	}
// // 	return user.To, nil

// // }
