package mock

import (
	"context"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type AreaStorageMock struct {
	mock.Mock
}

func (s *AreaStorageMock) CreateArea(ctx context.Context, area *entity.Area) error {
	args := s.Called(area)
	return args.Error(0)
}
