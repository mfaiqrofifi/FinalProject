package userRegister

import (
	"context"
	"social_media/drivers/databases/userStatus"
	"time"
)

type DomainRegister struct {
	ID        int
	Name      string
	Email     string
	Address   string
	Password  string
	Token     string
	Status    []userStatus.DBStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Usecase interface {
	Login(ctx context.Context, email string, password string) (DomainRegister, error)
	Register(ctx context.Context, name string, email string, address string, password string) (DomainRegister, error)
	DeteilRegister(ctx context.Context, IdUser int) ([]DomainRegister, error)
	GetFriend(ctx context.Context, email string) ([]DomainRegister, error)
}
type Repository interface {
	Login(ctx context.Context, email string, password string) (DomainRegister, error)
	Register(ctx context.Context, name string, email string, address string, password string) (DomainRegister, error)
	DeteilRegister(ctx context.Context, IdUser int) ([]DomainRegister, error)
	GetFriend(ctx context.Context, nama string) ([]DomainRegister, error)
}
