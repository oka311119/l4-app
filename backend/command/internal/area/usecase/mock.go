package usecase

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type AreaUseCaseMock struct {
	mock.Mock
}

func (m *AreaUseCaseMock) CreateDefaultArea(ctx context.Context) error {
	return args.Error(0)
}

func (m *AreaUseCaseMock) CreateArea(ctx context.Context, name string) error {
	args := m.Called(name)
	return args.Error(0)
}