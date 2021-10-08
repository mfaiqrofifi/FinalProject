package statusdesc

import (
	"context"
	"time"
)

type DetailStatus struct {
	ID        int
	Coment    string
	UserID    int
	Like      int
	Dislike   int
	StatusId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type CountLikeDisLike struct {
// 	Likes    int
// 	DisLikes int
// }

type Usecase interface {
	MakingComment(ctx context.Context, comment string, userId int, statusId int) (DetailStatus, error)
	MakingLike(ctx context.Context, userId int, statusId int) (DetailStatus, error)
	// HowManyLike(ctx context.Context, statusID int) ([]CountLikeDisLike, error)
}
type Repository interface {
	MakingComment(ctx context.Context, comment string, userId int, statusId int) (DetailStatus, error)
	MakingLike(ctx context.Context, userId int, statusId int) (DetailStatus, error)
	// HowManyLike(ctx context.Context, statusID int) ([]CountLikeDisLike, error)
}
