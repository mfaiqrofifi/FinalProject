package statusdesc

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUscase(repo Repository, timeOut time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repo,
		contextTimeout: timeOut,
	}
}

func (uc UserUsecase) MakingComment(ctx context.Context, comment string, userId int, statusId int) (DetailStatus, error) {
	if comment == "" {
		return DetailStatus{}, errors.New("comment empethy")
	}
	user, err := uc.Repo.MakingComment(ctx, comment, userId, statusId)
	if err != nil {
		return DetailStatus{}, err
	}
	return user, nil
}
func (uc UserUsecase) MakingLike(ctx context.Context, userId int, statusId int) (DetailStatus, error) {
	user, err := uc.Repo.MakingLike(ctx, userId, statusId)
	if err != nil {
		return DetailStatus{}, err
	}
	return user, nil
}

// func (uc UserUsecase) HowManyLike(ctx context.Context, statusID int) ([]CountLikeDisLike, error) {
// 	user, err := uc.Repo.HowManyLike(ctx, statusID)
// 	if err != nil {
// 		return []CountLikeDisLike{}, err
// 	}
// 	return user, nil
// }
