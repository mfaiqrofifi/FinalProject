package detailStatus

import (
	"context"
	"social_media/business/users/statusdesc"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) statusdesc.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) MakingLike(ctx context.Context, userId int, statusId int) (statusdesc.DetailStatus, error) {
	var user DetailStatusDB
	user.Like = userId
	user.UserID = userId
	user.DBStatusID = statusId
	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return statusdesc.DetailStatus{}, result.Error
	}
	return user.ToDomainComment(), nil
}

func (rep *MysqlUserRepository) MakingComment(ctx context.Context, comment string, userId int, statusId int) (statusdesc.DetailStatus, error) {
	var user DetailStatusDB
	user.Coment = comment
	user.UserID = userId
	user.DBStatusID = statusId
	result := rep.Conn.Create(&user)
	if result.Error != nil {
		return statusdesc.DetailStatus{}, result.Error
	}
	return user.ToDomainComment(), nil
}

// func (rep *MysqlUserRepository) HowManyLike()
