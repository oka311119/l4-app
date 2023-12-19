package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/auth/repository/mock"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "pepper", []byte("secret"), 86400)

	var (
		username = "user"
		password = "pass"
		
		ctx = context.Background()
		
		user = &entity.User{
			Username: username,
			Password: "c8b2505b76926abdc733523caa9f439142f66aa7293a7baaac0aed41a191eef6",	// sha256 of pass+salt+pepper
		}
	)

	// Sign Up
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, username, password)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}