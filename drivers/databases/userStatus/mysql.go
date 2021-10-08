package userStatus

import (
	"context"
	"social_media/business/users/userStatus"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) userStatus.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) CreateStatus(ctx context.Context, status string, IdUser int) (userStatus.UserStatus, error) {
	var user DBStatus
	user.Status = status
	user.DBRegistersID = IdUser
	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return userStatus.UserStatus{}, result.Error
	}
	return user.ToDomain(), nil

}

func (rep *MysqlUserRepository) SeeComment(ctx context.Context, Idstatus int) ([]userStatus.UserStatus, error) {
	var user []DBStatus
	resuult := rep.Conn.Preload("Coment").First(&user, "id=?", Idstatus)
	if resuult.Error != nil {
		return []userStatus.UserStatus{}, resuult.Error
	}
	return ToListDeteil(user), nil
}
