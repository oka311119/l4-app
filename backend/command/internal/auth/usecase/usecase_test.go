package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/auth/repository/mock"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/helpers"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "pepper", []byte("secret"), 86400, &helpers.MockSaltGenerator{})

	var (
		username = "user"
		password = "pass"
		salt, _ = uc.saltGen.Generate()
		
		ctx = context.Background()
		
		user = &entity.User {
			Username: username,
			Password: "ac1567a30817eae0e0b4ec52474e6be34469db8b59a09aa8a675518b01e7e547",	// sha256 of pass+salt+pepper
			Salt: salt,
		}
	)

	// Sign Up
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, username, password)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username).Return(user, nil)
	token, err := uc.SignIn(ctx, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}