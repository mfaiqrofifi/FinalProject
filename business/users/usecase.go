package users

// import (
// 	"context"
// 	"errors"
// 	middlewares "social_media/app/routes/middleware"
// 	"social_media/helpers/encrypt"
// 	"time"
// )

// type UserUsecase struct {
// 	ConfigJWT      middlewares.ConfigJWT
// 	Repo           Repository
// 	contextTimeout time.Duration
// }

// func NewUserUscase(repo Repository, timeOut time.Duration, configJWT middlewares.ConfigJWT) Usecase {
// 	return &UserUsecase{
// 		ConfigJWT:      configJWT,
// 		Repo:           repo,
// 		contextTimeout: timeOut,
// 	}
// }

// func (uc UserUsecase) Login(ctx context.Context, email string, password string) (DomainRegister, error) {
// 	if email == "" {
// 		return DomainRegister{}, errors.New("email empethy")
// 	}
// 	if password == "" {
// 		return DomainRegister{}, errors.New("email empethy")
// 	}
// 	user, err := uc.Repo.Login(ctx, email, password)
// 	temp := encrypt.ValidateHash(password, user.Password)
// 	if !temp {
// 		return DomainRegister{}, errors.New("password wrong")
// 	}
// 	if err != nil {
// 		return DomainRegister{}, err
// 	}
// 	user.Token, err = uc.ConfigJWT.GenerateToken(user.ID, "user")
// 	if err != nil {
// 		return DomainRegister{}, err
// 	}
// 	return user, nil
// }

// func (uc UserUsecase) Register(ctx context.Context, name string, email string, address string, password string) (DomainRegister, error) {
// 	if name == "" {
// 		return DomainRegister{}, errors.New("name empethy")
// 	}
// 	if email == "" {
// 		return DomainRegister{}, errors.New("email empethy")
// 	}
// 	if password == "" {
// 		return DomainRegister{}, errors.New("email password")
// 	}
// 	var errPass error
// 	password, errPass = encrypt.Hash(password)
// 	if errPass != nil {
// 		return DomainRegister{}, errPass
// 	}
// 	user, err := uc.Repo.Register(ctx, name, email, address, password)
// 	if err != nil {
// 		return DomainRegister{}, err
// 	}
// 	return user, nil
// }

// func (uc UserUsecase) CreateStatus(ctx context.Context, status string, IdUser int) (UserStatus, error) {
// 	if status == "" {
// 		return UserStatus{}, errors.New("satus empethy")
// 	}
// 	user, err := uc.Repo.CreateStatus(ctx, status, IdUser)
// 	if err != nil {
// 		return UserStatus{}, err
// 	}
// 	return user, nil
// }

// // func (uc UserUsecase) UserDeteil(ctx context.Context, name string, address string, status []UserStatus) (DomainRegister, error) {
// // 	user, err := uc.Repo.UserDeteil(ctx, name, address, status)
// // 	if err != nil {
// // 		return DomainRegister{}, err
// // 	}
// // 	return user, nil
// // }
