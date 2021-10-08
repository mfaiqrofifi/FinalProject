package detailStatus

import (
	"social_media/business/users/statusdesc"
	"time"
)

type DetailStatusDB struct {
	ID         int `gorm:"primaryKey"`
	Coment     string
	UserID     int
	Like       int
	Dislike    int
	DBStatusID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type CountLikeDisLikeDB struct {
	Like    int
	DisLike int
}

// func (uc UserUsecase) CountLikeDisLike(ctx context.Context, statusID int) ([]DomainRegister, error) {
// 	user, err := uc.Repo.DeteilRegister(ctx, IdUser)
// 	if err != nil {
// 		return []DomainRegister{}, err
// 	}
// 	return user, nil
// }

func (user *DetailStatusDB) ToDomainComment() statusdesc.DetailStatus {
	return statusdesc.DetailStatus{
		ID:        user.ID,
		Coment:    user.Coment,
		UserID:    user.UserID,
		Like:      user.Like,
		Dislike:   user.Dislike,
		StatusId:  user.DBStatusID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func FromDomainComment(domain statusdesc.DetailStatus) DetailStatusDB {
	return DetailStatusDB{
		ID:         domain.ID,
		Coment:     domain.Coment,
		UserID:     domain.UserID,
		Like:       domain.Like,
		Dislike:    domain.Dislike,
		DBStatusID: domain.StatusId,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
