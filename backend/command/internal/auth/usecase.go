package auth

import "github.com/oka311119/l4-app/backend/command/internal/handler"

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (string, error)
	ParseToken(ctx context.Context, user *models.User) (*models.User, error)
	DeleteAccount(ctx context.Context, user *models.User) error
}