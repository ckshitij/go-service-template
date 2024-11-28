package users

import (
	"context"
)

type UsersRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UsersService interface {
	CreateUser(ctx context.Context, req ServiceRequest) (*User, error)
	GetUser(ctx context.Context, req ServiceRequest) (*User, error)
}
