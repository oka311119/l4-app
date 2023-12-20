package auth

import (
	"context"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
}
