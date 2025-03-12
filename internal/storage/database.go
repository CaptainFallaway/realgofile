package storage

import (
	"context"
)

type Repository interface {
	InsertUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, username string, user User) error
	DeleteUser(ctx context.Context, uid string) error
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByUid(ctx context.Context, uid string) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
}
