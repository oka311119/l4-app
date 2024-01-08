package usecase

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type AreaUseCaseMock struct {
	mock.Mock
}

func (m *AreaUseCaseMock) CreateArea(ctx context.Context, userID, name string) error {
	args := m.Called(userID, name)
	return args.Error(0)
}
