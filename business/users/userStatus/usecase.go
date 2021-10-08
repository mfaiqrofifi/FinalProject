package userStatus

import (
	"context"
	"errors"
	middlewares "social_media/app/routes/middleware"
	"time"
)

type UserUsecase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUscase(repo Repository, timeOut time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeOut,
	}
}

func (uc UserUsecase) CreateStatus(ctx context.Context, status string, IdUser int) (UserStatus, error) {
	if status == "" {
		return UserStatus{}, errors.New("satus empethy")
	}
	user, err := uc.Repo.CreateStatus(ctx, status, IdUser)
	if err != nil {
		return UserStatus{}, err
	}
	return user, nil
}
func (uc UserUsecase) SeeComment(ctx context.Context, Idstatus int) ([]UserStatus, error) {
	user, err := uc.Repo.SeeComment(ctx, Idstatus)
	if err != nil {
		return []UserStatus{}, err
	}
	return user, nil
}
