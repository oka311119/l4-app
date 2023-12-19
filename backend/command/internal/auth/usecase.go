package auth

import (
	"context"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entity.User, error)
}