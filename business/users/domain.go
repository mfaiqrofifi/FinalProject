package users

// import (
// 	"context"
// 	"time"
// )

// type DomainRegister struct {
// 	ID        int
// 	Name      string
// 	Email     string
// 	Address   string
// 	Password  string
// 	Token     string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }
// type UserDeteil struct {
// 	DomainRegister DomainRegister
// 	Status         []UserStatus
// }

// type UserStatus struct {
// 	ID        int
// 	Status    string
// 	IdUser    int
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// type Usecase interface {
// 	Login(ctx context.Context, email string, password string) (DomainRegister, error)
// 	Register(ctx context.Context, name string, email string, address string, password string) (DomainRegister, error)
// 	CreateStatus(ctx context.Context, status string, IdUser int) (UserStatus, error)
// 	// UserDeteil(ctx context.Context, name string, address string, status []UserStatus) (DomainRegister, error)
// }

// type Repository interface {
// 	Login(ctx context.Context, email string, password string) (DomainRegister, error)
// 	Register(ctx context.Context, name string, email string, address string, password string) (DomainRegister, error)
// 	CreateStatus(ctx context.Context, status string, IdUser int) (UserStatus, error)
// 	// UserDeteil(ctx context.Context, name string, address string, status []UserStatus) (DomainRegister, error)
// }
