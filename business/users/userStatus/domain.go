package userStatus

import (
	"context"
	"social_media/drivers/databases/detailStatus"
	"time"
)

type UserStatus struct {
	ID        int
	Status    string
	IdUser    int
	Comment   []detailStatus.DetailStatusDB
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	CreateStatus(ctx context.Context, status string, IdUser int) (UserStatus, error)
	SeeComment(ctx context.Context, Idstatus int) ([]UserStatus, error)
}

type Repository interface {
	CreateStatus(ctx context.Context, status string, IdUser int) (UserStatus, error)
	SeeComment(ctx context.Context, Idstatus int) ([]UserStatus, error)
}
