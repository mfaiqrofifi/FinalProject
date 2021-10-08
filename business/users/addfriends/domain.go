package addfriends

import "context"

type AddFriends struct {
	ID       int
	UserID   int
	Name     string
	Email    string
	Address  string
	FriendID string
}

type UserUsecase interface {
	AddFriends(ctx context.Context, userid int, friendId int)
}

type Repository interface {
	AddFriends(ctx context.Context, userid int, friendId int)
}
