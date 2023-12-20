package mock

import (
	"context"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func(s *UserStorageMock) CreateUser(ctx context.Context, user *entity.User) error {
	args := s.Called(user)
	return args.Error(0)
}

func (s *UserStorageMock) GetUser(ctx context.Context, username string) (*entity.User, error) {
	args := s.Called(username)
	return args.Get(0).(*entity.User), args.Error(1)
}