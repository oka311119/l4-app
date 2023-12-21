package usecase

import (
	"context"

	"github.com/oka311119/l4-app/backend/command/internal/area"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

type AreaUseCase struct {
	areaRepo area.Repository
}

func NewAreaUseCase(
	areaRepo area.Repository,
) *AreaUseCase {
	return &AreaUseCase{
		areaRepo: areaRepo,
	}
}

func (a *AreaUseCase) CreateDefaultArea(ctx context.Context) error {
	area := &entity.Area{
		ID
	}

	return a.areaRepo.CreateArea(ctx, area)
}