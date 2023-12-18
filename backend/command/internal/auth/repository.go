package auth

import "github.com/oka311119/l4-app/backend/command/internal/handler"

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
}