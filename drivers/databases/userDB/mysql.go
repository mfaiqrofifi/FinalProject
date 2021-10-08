package userDB

import (
	"context"
	"social_media/business/users/userRegister"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) userRegister.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (userRegister.DomainRegister, error) {
	var user DBRegisters
	result := rep.Conn.First(&user, "email=?", email)
	if result.Error != nil {
		return userRegister.DomainRegister{}, result.Error
	}
	return user.ToDomain(), nil

}
func (rep *MysqlUserRepository) Register(ctx context.Context, name string, email string, address string, password string) (userRegister.DomainRegister, error) {
	var user DBRegisters
	user.Nama = name
	user.Address = address
	user.Email = email
	user.Password = password
	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return userRegister.DomainRegister{}, result.Error
	}
	return user.ToDomain(), nil

}

func (rep *MysqlUserRepository) DeteilRegister(ctx context.Context, IdUser int) ([]userRegister.DomainRegister, error) {
	var user []DBRegisters
	result := rep.Conn.Debug().Preload(clause.Associations).First(&user, "id=?", IdUser)
	if result.Error != nil {
		return []userRegister.DomainRegister{}, result.Error
	}
	return ToListDeteil(user), nil
}
func (rep *MysqlUserRepository) GetFriend(ctx context.Context, nama string) ([]userRegister.DomainRegister, error) {
	var user []DBRegisters
	result := rep.Conn.Where("nama Like ?", "%"+nama+"%").Find(&user)
	if result.Error != nil {
		return []userRegister.DomainRegister{}, result.Error
	}
	return ToListDeteil(user), nil
}
