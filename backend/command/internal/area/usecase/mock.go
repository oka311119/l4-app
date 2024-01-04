package usecase

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type AreaUseCaseMock struct {
	mock.Mock
}

func (m *AreaUseCaseMock) CreateDefaultArea(ctx context.Context, userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

func (m *AreaUseCaseMock) CreateArea(ctx context.Context, userID, name string) error {
	args := m.Called(userID, name)
	return args.Error(0)
}
