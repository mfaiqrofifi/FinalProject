package userRegister

import (
	"context"
	"errors"
	middlewares "social_media/app/routes/middleware"
	"social_media/helpers/encrypt"

	// "social_media/helpers/encrypt"
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

func (uc UserUsecase) Login(ctx context.Context, email string, password string) (DomainRegister, error) {
	if email == "" {
		return DomainRegister{}, errors.New("email empethy")
	}
	if password == "" {
		return DomainRegister{}, errors.New("password empethy")
	}
	user, err := uc.Repo.Login(ctx, email, password)
	temp := encrypt.ValidateHash(password, user.Password)
	if !temp {
		return DomainRegister{}, errors.New("password wrong")
	}
	if err != nil {
		return DomainRegister{}, err
	}
	user.Token, err = uc.ConfigJWT.GenerateToken(user.ID, "user")
	if err != nil {
		return DomainRegister{}, err
	}
	return user, nil
}

func (uc UserUsecase) Register(ctx context.Context, name string, email string, address string, password string) (DomainRegister, error) {
	if name == "" {
		return DomainRegister{}, errors.New("name empethy")
	}
	if email == "" {
		return DomainRegister{}, errors.New("email empethy")
	}
	if password == "" {
		return DomainRegister{}, errors.New("email password")
	}
	var errPass error
	password, errPass = encrypt.Hash(password)
	if errPass != nil {
		return DomainRegister{}, errPass
	}
	user, err := uc.Repo.Register(ctx, name, email, address, password)
	if err != nil {
		return DomainRegister{}, err
	}
	return user, nil
}
func (uc UserUsecase) DeteilRegister(ctx context.Context, IdUser int) ([]DomainRegister, error) {
	user, err := uc.Repo.DeteilRegister(ctx, IdUser)
	if err != nil {
		return []DomainRegister{}, err
	}
	return user, nil
}
func (uc UserUsecase) GetFriend(ctx context.Context, nama string) ([]DomainRegister, error) {
	if nama != "" {
		return []DomainRegister{}, errors.New("enter the name")
	}
	user, err := uc.Repo.GetFriend(ctx, nama)
	if err != nil {
		return []DomainRegister{}, err
	}
	return user, nil
}
